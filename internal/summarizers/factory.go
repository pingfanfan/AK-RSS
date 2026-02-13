package summarizers

import (
	"fmt"

	"github.com/opmlwatch/opmlwatch/internal/config"
	"github.com/opmlwatch/opmlwatch/internal/summarizers/none"
	"github.com/opmlwatch/opmlwatch/internal/summarizers/openai"
)

func BuildFromConfig(cfg config.SummarizerConfig) (Summarizer, error) {
	if !cfg.Enabled || cfg.Type == "" || cfg.Type == "none" {
		return none.New(), nil
	}
	switch cfg.Type {
	case "openai":
		return openai.New(cfg.Config)
	default:
		return nil, fmt.Errorf("unsupported summarizer type: %s", cfg.Type)
	}
}
