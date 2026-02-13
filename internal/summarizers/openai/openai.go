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
	baseURL string
	apiKey  string
	model   string
	client  *http.Client
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
	baseURL := strings.TrimSpace(cfg["base_url"])
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	apiKey := strings.TrimSpace(cfg["api_key"])
	if strings.HasPrefix(apiKey, "env:") {
		apiKey = os.Getenv(strings.TrimPrefix(apiKey, "env:"))
	}
	if apiKey == "" {
		return nil, fmt.Errorf("openai summarizer missing api_key")
	}
	model := strings.TrimSpace(cfg["model"])
	if model == "" {
		model = "gpt-4o-mini"
	}

	return &Summarizer{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		apiKey:  apiKey,
		model:   model,
		client:  &http.Client{Timeout: 20 * time.Second},
	}, nil
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

	prompt := "Analyze this feed update and return exactly 3 lines in Chinese: WHAT: ..., WHY: ..., ACTION: ...\\n" +
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
	what, why, action := parseThreeLines(content)
	return core.Analysis{
		What:   what,
		Why:    why,
		Action: action,
		Model:  s.model,
	}, nil
}

func parseThreeLines(s string) (what, why, action string) {
	lines := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	for _, line := range lines {
		l := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(strings.ToUpper(l), "WHAT:"):
			what = strings.TrimSpace(l[5:])
		case strings.HasPrefix(strings.ToUpper(l), "WHY:"):
			why = strings.TrimSpace(l[4:])
		case strings.HasPrefix(strings.ToUpper(l), "ACTION:"):
			action = strings.TrimSpace(l[7:])
		}
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
	return
}
