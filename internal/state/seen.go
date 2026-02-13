package state

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type SeenStore struct {
	path string
	seen map[string]int64
}

func New(path string) (*SeenStore, error) {
	if strings.TrimSpace(path) == "" {
		path = ".opmlwatch/seen_links.json"
	}
	s := &SeenStore{path: path, seen: map[string]int64{}}
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SeenStore) FilterNew(entries []core.Entry) []core.Entry {
	out := make([]core.Entry, 0, len(entries))
	for _, e := range entries {
		k := key(e)
		if _, ok := s.seen[k]; ok {
			continue
		}
		out = append(out, e)
	}
	return out
}

func (s *SeenStore) MarkSeen(entries []core.Entry) {
	now := time.Now().Unix()
	for _, e := range entries {
		s.seen[key(e)] = now
	}
}

func (s *SeenStore) Save() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(s.seen, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

func (s *SeenStore) load() error {
	b, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.seen)
}

func key(e core.Entry) string {
	if strings.TrimSpace(e.Link) != "" {
		return strings.ToLower(strings.TrimSpace(e.Link))
	}
	return strings.ToLower(strings.TrimSpace(e.Title))
}
