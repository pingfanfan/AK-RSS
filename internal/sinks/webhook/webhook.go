package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Sink struct {
	name   string
	url    string
	dryRun bool
	client *http.Client
}

func New(name string, cfg map[string]string, dryRun bool) (*Sink, error) {
	u := cfg["url"]
	if u == "" {
		return nil, fmt.Errorf("webhook sink %s missing config.url", name)
	}
	return &Sink{
		name:   name,
		url:    u,
		dryRun: dryRun,
		client: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (s *Sink) Name() string { return s.name }

func (s *Sink) Send(ctx context.Context, batch core.Batch) error {
	if s.dryRun {
		return nil
	}
	b, err := json.Marshal(batch)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("webhook %s returned %s", s.name, resp.Status)
	}
	return nil
}
