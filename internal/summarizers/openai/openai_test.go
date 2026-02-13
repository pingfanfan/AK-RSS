package openai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

func TestParseFiveLines(t *testing.T) {
	tldr, what, why, action, tweet := parseFiveLines("TLDR: T\nWHAT: A\nWHY: B\nACTION: C\nTWEET: X")
	if tldr != "T" || what != "A" || why != "B" || action != "C" || tweet != "X" {
		t.Fatalf("unexpected parse result: %q %q %q %q %q", tldr, what, why, action, tweet)
	}
}

func TestSummarize(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST")
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"choices": []map[string]any{{
				"message": map[string]any{"content": "TLDR: T\nWHAT: W\nWHY: Y\nACTION: A\nTWEET: X"},
			}},
		})
	}))
	defer ts.Close()

	s, err := New(map[string]string{
		"base_url": ts.URL,
		"api_key":  "test-key",
		"model":    "gpt-4o-mini",
	})
	if err != nil {
		t.Fatalf("new summarizer: %v", err)
	}

	out, err := s.Summarize(context.Background(), core.Entry{Title: "hello", Link: "https://example.com", Content: "world"})
	if err != nil {
		t.Fatalf("summarize: %v", err)
	}
	if out.TLDR != "T" || out.What != "W" || out.Why != "Y" || out.Action != "A" || out.Tweet != "X" {
		t.Fatalf("unexpected summary: %+v", out)
	}
	if out.Model != "gpt-4o-mini" {
		t.Fatalf("unexpected model: %s", out.Model)
	}
}

func TestNewResolveEnv(t *testing.T) {
	_ = os.Setenv("AI_BASE_URL_TEST", "https://example.com/v1")
	_ = os.Setenv("AI_MODEL_TEST", "m")
	_ = os.Setenv("AI_KEY_TEST", "k")
	defer os.Unsetenv("AI_BASE_URL_TEST")
	defer os.Unsetenv("AI_MODEL_TEST")
	defer os.Unsetenv("AI_KEY_TEST")

	s, err := New(map[string]string{
		"base_url": "env:AI_BASE_URL_TEST",
		"model":    "env:AI_MODEL_TEST",
		"api_key":  "env:AI_KEY_TEST",
	})
	if err != nil {
		t.Fatalf("new summarizer: %v", err)
	}
	if s.baseURL != "https://example.com/v1" || s.model != "m" || s.apiKey != "k" {
		t.Fatalf("unexpected resolved values: %+v", s)
	}
}
