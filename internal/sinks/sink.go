package sinks

import (
	"context"
	"fmt"

	"github.com/opmlwatch/opmlwatch/internal/config"
	"github.com/opmlwatch/opmlwatch/internal/core"
	"github.com/opmlwatch/opmlwatch/internal/sinks/email"
	"github.com/opmlwatch/opmlwatch/internal/sinks/markdown"
	"github.com/opmlwatch/opmlwatch/internal/sinks/webhook"
)

type Sink interface {
	Name() string
	Send(ctx context.Context, batch core.Batch) error
}

func BuildFromConfig(cfgs []config.SinkConfig, dryRun bool) ([]Sink, error) {
	out := make([]Sink, 0, len(cfgs))
	for _, c := range cfgs {
		if !c.Enabled {
			continue
		}
		switch c.Type {
		case "webhook":
			s, err := webhook.New(c.Name, c.Config, dryRun)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		case "markdown":
			s, err := markdown.New(c.Name, c.Config)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		case "email":
			s, err := email.New(c.Name, c.Config)
			if err != nil {
				return nil, err
			}
			out = append(out, s)
		default:
			return nil, fmt.Errorf("unsupported sink type: %s", c.Type)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no enabled sinks")
	}
	return out, nil
}
