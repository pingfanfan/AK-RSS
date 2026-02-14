package site

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

func TestSiteSinkWritesFiles(t *testing.T) {
	d := t.TempDir()
	s, err := New("site", map[string]string{
		"updates_path": filepath.Join(d, "updates.json"),
		"daily_path":   filepath.Join(d, "daily.json"),
		"max_updates":  "10",
		"max_days":     "7",
		"timezone":     "Europe/London",
	})
	if err != nil {
		t.Fatalf("new sink: %v", err)
	}

	batch := core.Batch{
		RunAt: time.Now().UTC(),
		Entries: []core.Entry{
			{
				FeedTitle: "feed-a",
				Title:     "title-a",
				Link:      "https://example.com/a",
				Published: time.Now().UTC(),
				Analysis: &core.Analysis{
					TLDR:  "tldr",
					Tweet: "tweet",
				},
			},
		},
	}
	if err := s.Send(context.Background(), batch); err != nil {
		t.Fatalf("send: %v", err)
	}

	if _, err := os.Stat(filepath.Join(d, "updates.json")); err != nil {
		t.Fatalf("updates file missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(d, "daily.json")); err != nil {
		t.Fatalf("daily file missing: %v", err)
	}
}
