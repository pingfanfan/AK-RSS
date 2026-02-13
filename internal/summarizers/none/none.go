package none

import (
	"context"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Summarizer struct{}

func New() *Summarizer { return &Summarizer{} }

func (s *Summarizer) Name() string { return "none" }

func (s *Summarizer) Summarize(_ context.Context, entry core.Entry) (core.Analysis, error) {
	return core.Analysis{
		What:   entry.Title,
		Why:    "Fresh update from monitored feed.",
		Action: "Open source link and decide whether to read now or add to digest backlog.",
		Model:  "none",
	}, nil
}
