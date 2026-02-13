package summarizers

import (
	"context"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Summarizer interface {
	Name() string
	Summarize(ctx context.Context, entry core.Entry) (core.Analysis, error)
}
