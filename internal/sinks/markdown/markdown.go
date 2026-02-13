package markdown

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Sink struct {
	name string
	path string
}

func New(name string, cfg map[string]string) (*Sink, error) {
	p := cfg["path"]
	if p == "" {
		return nil, fmt.Errorf("markdown sink %s missing config.path", name)
	}
	return &Sink{name: name, path: p}, nil
}

func (s *Sink) Name() string { return s.name }

func (s *Sink) Send(_ context.Context, batch core.Batch) error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	var b strings.Builder
	b.WriteString("# OPMLWatch Digest\n\n")
	b.WriteString("Generated at: " + batch.RunAt.Format("2006-01-02 15:04:05 MST") + "\n\n")
	for _, e := range batch.Entries {
		b.WriteString("- [" + e.Title + "](" + e.Link + ")")
		if e.FeedTitle != "" {
			b.WriteString(" (`" + e.FeedTitle + "`)")
		}
		b.WriteString("\n")
		if e.Analysis != nil {
			b.WriteString("  - WHAT: " + e.Analysis.What + "\n")
			b.WriteString("  - WHY: " + e.Analysis.Why + "\n")
			b.WriteString("  - ACTION: " + e.Analysis.Action + "\n")
		}
	}
	return os.WriteFile(s.path, []byte(b.String()), 0o644)
}
