package state

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

func TestSeenStore(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "seen.json")
	s, err := New(p)
	if err != nil {
		t.Fatalf("new store: %v", err)
	}
	entries := []core.Entry{{Title: "A", Link: "https://example.com/a"}}
	if got := s.FilterNew(entries); len(got) != 1 {
		t.Fatalf("expected new entries")
	}
	s.MarkSeen(entries)
	if got := s.FilterNew(entries); len(got) != 0 {
		t.Fatalf("expected no new entries")
	}
	if err := s.Save(); err != nil {
		t.Fatalf("save: %v", err)
	}
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("stat: %v", err)
	}
}
