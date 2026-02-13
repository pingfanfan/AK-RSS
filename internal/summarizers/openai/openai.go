package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Summarizer struct {
	baseURL       string
	apiKey        string
	model         string
	readerProfile string
	client        *http.Client
}

type chatRequest struct {
	Model       string        `json:"model"`
	Messages    []chatMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func New(cfg map[string]string) (*Summarizer, error) {
	baseURL := resolve(cfg["base_url"])
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	apiKey := resolve(cfg["api_key"])
	if apiKey == "" {
		return nil, fmt.Errorf("openai summarizer missing api_key")
	}
	model := resolve(cfg["model"])
	if model == "" {
		model = "gpt-4o-mini"
	}
	readerProfile := resolve(cfg["reader_profile"])
	if readerProfile == "" {
		readerProfile = "中文读者，关注技术趋势与可执行建议。"
	}

	return &Summarizer{
		baseURL:       strings.TrimSuffix(baseURL, "/"),
		apiKey:        apiKey,
		model:         model,
		readerProfile: readerProfile,
		client:        &http.Client{Timeout: 20 * time.Second},
	}, nil
}

func resolve(v string) string {
	v = strings.TrimSpace(v)
	if strings.HasPrefix(v, "env:") {
		return strings.TrimSpace(os.Getenv(strings.TrimPrefix(v, "env:")))
	}
	return v
}

func (s *Summarizer) Name() string { return "openai" }

func (s *Summarizer) Summarize(ctx context.Context, entry core.Entry) (core.Analysis, error) {
	text := strings.TrimSpace(entry.Content)
	if text == "" {
		text = entry.Title
	}
	if len(text) > 3000 {
		text = text[:3000]
	}

	prompt := "请基于读者画像，分析这篇最新博客更新，并严格返回 5 行：\n" +
		"TLDR: ...\n" +
		"WHAT: ...\n" +
		"WHY: ...\n" +
		"ACTION: ...\n" +
		"TWEET: ...\n\n" +
		"要求：\n" +
		"- 使用简体中文\n" +
		"- 内容具体可执行，避免空话\n" +
		"- TWEET 控制在 260 字以内，可直接发布\n\n" +
		"读者画像: " + s.readerProfile + "\n\n" +
		"Title: " + entry.Title + "\\n" +
		"Link: " + entry.Link + "\\n" +
		"Content: " + text

	reqBody := chatRequest{
		Model: s.model,
		Messages: []chatMessage{
			{Role: "system", Content: "You are a concise analyst."},
			{Role: "user", Content: prompt},
		},
		Temperature: 0.2,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return core.Analysis{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+"/chat/completions", bytes.NewReader(b))
	if err != nil {
		return core.Analysis{}, err
	}
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return core.Analysis{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return core.Analysis{}, fmt.Errorf("openai status: %s", resp.Status)
	}

	var out chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return core.Analysis{}, err
	}
	if len(out.Choices) == 0 {
		return core.Analysis{}, fmt.Errorf("openai empty choices")
	}

	content := out.Choices[0].Message.Content
	tldr, what, why, action, tweet := parseFiveLines(content)
	return core.Analysis{
		TLDR:   tldr,
		What:   what,
		Why:    why,
		Action: action,
		Tweet:  tweet,
		Model:  s.model,
	}, nil
}

func parseFiveLines(s string) (tldr, what, why, action, tweet string) {
	lines := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	for _, line := range lines {
		l := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(strings.ToUpper(l), "TLDR:"):
			tldr = strings.TrimSpace(l[5:])
		case strings.HasPrefix(strings.ToUpper(l), "WHAT:"):
			what = strings.TrimSpace(l[5:])
		case strings.HasPrefix(strings.ToUpper(l), "WHY:"):
			why = strings.TrimSpace(l[4:])
		case strings.HasPrefix(strings.ToUpper(l), "ACTION:"):
			action = strings.TrimSpace(l[7:])
		case strings.HasPrefix(strings.ToUpper(l), "TWEET:"):
			tweet = strings.TrimSpace(l[6:])
		}
	}
	if tldr == "" {
		tldr = "一条值得关注的博客更新。"
	}
	if what == "" {
		what = "新条目更新"
	}
	if why == "" {
		why = "信息可能影响你的关注主题"
	}
	if action == "" {
		action = "打开原文快速判断是否加入待办"
	}
	if tweet == "" {
		tweet = what
	}
	return
}
