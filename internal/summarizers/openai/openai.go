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
	"unicode/utf8"

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

type structuredOutput struct {
	TitleZh string `json:"title_zh"`
	TitleEn string `json:"title_en"`

	TLDRZh string `json:"tldr_zh"`
	TLDREn string `json:"tldr_en"`
	WhatZh string `json:"what_zh"`
	WhatEn string `json:"what_en"`
	WhyZh  string `json:"why_zh"`
	WhyEn  string `json:"why_en"`

	ActionZh string `json:"action_zh"`
	ActionEn string `json:"action_en"`
	TweetZh  string `json:"tweet_zh"`
	TweetEn  string `json:"tweet_en"`

	XZh        string `json:"x_zh"`
	XEn        string `json:"x_en"`
	LinkedInZh string `json:"linkedin_zh"`
	LinkedInEn string `json:"linkedin_en"`
	ThreadsZh  string `json:"threads_zh"`
	ThreadsEn  string `json:"threads_en"`
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
	if len(text) > 2800 {
		text = text[:2800]
	}

	prompt := "" +
		"请基于读者画像，输出一份可传播的双语分析。\n" +
		"严格只返回 JSON 对象，不要 markdown、不要解释、不要代码块。\n" +
		"字段必须完整：\n" +
		"{\n" +
		"  \"title_zh\":\"...\",\n" +
		"  \"title_en\":\"...\",\n" +
		"  \"tldr_zh\":\"...\",\n" +
		"  \"tldr_en\":\"...\",\n" +
		"  \"what_zh\":\"...\",\n" +
		"  \"what_en\":\"...\",\n" +
		"  \"why_zh\":\"...\",\n" +
		"  \"why_en\":\"...\",\n" +
		"  \"action_zh\":\"...\",\n" +
		"  \"action_en\":\"...\",\n" +
		"  \"tweet_zh\":\"...\",\n" +
		"  \"tweet_en\":\"...\",\n" +
		"  \"x_zh\":\"...\",\n" +
		"  \"x_en\":\"...\",\n" +
		"  \"linkedin_zh\":\"...\",\n" +
		"  \"linkedin_en\":\"...\",\n" +
		"  \"threads_zh\":\"...\",\n" +
		"  \"threads_en\":\"...\"\n" +
		"}\n\n" +
		"生成规则：\n" +
		"- 中文字段用简体中文，英文字段用自然英文。\n" +
		"- title_* 是标题翻译，保留原意。\n" +
		"- tldr/what/why/action 要具体，面向技术读者。\n" +
		"- tweet_* 与 x_* 为短文案，<=260 字符，可直接发。\n" +
		"- linkedin_* 为职业平台语气，3~5 行，含行动建议。\n" +
		"- threads_* 为更口语的讨论引导语气，2~4 行。\n" +
		"- 避免编造事实，基于给定内容。\n\n" +
		"读者画像: " + s.readerProfile + "\n\n" +
		"Title: " + entry.Title + "\n" +
		"Link: " + entry.Link + "\n" +
		"Content: " + text

	reqBody := chatRequest{
		Model: s.model,
		Messages: []chatMessage{
			{Role: "system", Content: "You are a precise analyst that returns strict JSON when asked."},
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

	content := strings.TrimSpace(out.Choices[0].Message.Content)
	if parsed, ok := parseStructured(content); ok {
		a := toAnalysis(parsed, entry)
		a.Model = s.model
		return a, nil
	}

	// Fallback for older/loose model outputs.
	tldr, what, why, action, tweet := parseFiveLines(content)
	return core.Analysis{
		TLDR:       tldr,
		What:       what,
		Why:        why,
		Action:     action,
		Tweet:      tweet,
		TitleZh:    entry.Title,
		TitleEn:    entry.Title,
		TLDRZh:     tldr,
		TLDREn:     tldr,
		WhatZh:     what,
		WhatEn:     what,
		WhyZh:      why,
		WhyEn:      why,
		ActionZh:   action,
		ActionEn:   action,
		TweetZh:    tweet,
		TweetEn:    tweet,
		XZh:        tweet,
		XEn:        tweet,
		LinkedInZh: tweet,
		LinkedInEn: tweet,
		ThreadsZh:  tweet,
		ThreadsEn:  tweet,
		Model:      s.model,
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

func parseStructured(raw string) (structuredOutput, bool) {
	var out structuredOutput
	candidate := strings.TrimSpace(raw)
	if candidate == "" {
		return out, false
	}
	if err := json.Unmarshal([]byte(candidate), &out); err == nil {
		return out, true
	}
	obj := extractJSONObject(candidate)
	if obj == "" {
		return out, false
	}
	if err := json.Unmarshal([]byte(obj), &out); err != nil {
		return out, false
	}
	return out, true
}

func extractJSONObject(s string) string {
	start := strings.Index(s, "{")
	if start < 0 {
		return ""
	}
	depth := 0
	inString := false
	escaped := false
	for i := start; i < len(s); i++ {
		ch := s[i]
		if inString {
			if escaped {
				escaped = false
				continue
			}
			if ch == '\\' {
				escaped = true
				continue
			}
			if ch == '"' {
				inString = false
			}
			continue
		}
		switch ch {
		case '"':
			inString = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[start : i+1]
			}
		}
	}
	return ""
}

func toAnalysis(o structuredOutput, e core.Entry) core.Analysis {
	titleZh := firstNonEmpty(o.TitleZh, e.Title)
	titleEn := firstNonEmpty(o.TitleEn, e.Title)

	tldrZh := firstNonEmpty(o.TLDRZh, o.WhatZh, o.TweetZh, e.Title)
	tldrEn := firstNonEmpty(o.TLDREn, o.WhatEn, o.TweetEn, e.Title)
	whatZh := firstNonEmpty(o.WhatZh, tldrZh)
	whatEn := firstNonEmpty(o.WhatEn, tldrEn)
	whyZh := firstNonEmpty(o.WhyZh, "这条更新可能影响你的技术判断。")
	whyEn := firstNonEmpty(o.WhyEn, "This update can affect your technical decisions.")
	actionZh := firstNonEmpty(o.ActionZh, "打开原文并判断是否进入你的行动清单。")
	actionEn := firstNonEmpty(o.ActionEn, "Open the source and decide if it belongs in your action queue.")

	tweetZh := clipRunes(firstNonEmpty(o.TweetZh, o.XZh, whatZh+" "+e.Link), 260)
	tweetEn := clipRunes(firstNonEmpty(o.TweetEn, o.XEn, whatEn+" "+e.Link), 260)

	xZh := clipRunes(firstNonEmpty(o.XZh, tweetZh), 260)
	xEn := clipRunes(firstNonEmpty(o.XEn, tweetEn), 260)

	linkedInZh := firstNonEmpty(o.LinkedInZh, tweetZh)
	linkedInEn := firstNonEmpty(o.LinkedInEn, tweetEn)
	threadsZh := firstNonEmpty(o.ThreadsZh, tweetZh)
	threadsEn := firstNonEmpty(o.ThreadsEn, tweetEn)

	return core.Analysis{
		TLDR:   firstNonEmpty(tldrZh, tldrEn),
		What:   firstNonEmpty(whatZh, whatEn),
		Why:    firstNonEmpty(whyZh, whyEn),
		Action: firstNonEmpty(actionZh, actionEn),
		Tweet:  firstNonEmpty(xZh, tweetZh, xEn, tweetEn),

		TitleZh:  titleZh,
		TitleEn:  titleEn,
		TLDRZh:   tldrZh,
		TLDREn:   tldrEn,
		WhatZh:   whatZh,
		WhatEn:   whatEn,
		WhyZh:    whyZh,
		WhyEn:    whyEn,
		ActionZh: actionZh,
		ActionEn: actionEn,
		TweetZh:  tweetZh,
		TweetEn:  tweetEn,

		XZh:        xZh,
		XEn:        xEn,
		LinkedInZh: linkedInZh,
		LinkedInEn: linkedInEn,
		ThreadsZh:  threadsZh,
		ThreadsEn:  threadsEn,
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func clipRunes(s string, max int) string {
	if max <= 0 {
		return ""
	}
	s = strings.TrimSpace(s)
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	r := []rune(s)
	if len(r) <= max {
		return string(r)
	}
	return strings.TrimSpace(string(r[:max]))
}
