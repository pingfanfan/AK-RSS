package email

import (
	"strings"
	"testing"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

func TestSplitCSV(t *testing.T) {
	out := splitCSV("a@example.com, b@example.com ,,")
	if len(out) != 2 {
		t.Fatalf("expected 2, got %d", len(out))
	}
}

func TestBuildBodyIncludesAnalysis(t *testing.T) {
	batch := core.Batch{
		RunAt: time.Unix(0, 0).UTC(),
		Entries: []core.Entry{{
			Title: "T",
			Link:  "https://example.com",
			Analysis: &core.Analysis{
				What:   "W",
				Why:    "Y",
				Action: "A",
			},
		}},
	}
	body := buildBody(batch)
	for _, k := range []string{"WHAT: W", "WHY: Y", "ACTION: A"} {
		if !strings.Contains(body, k) {
			t.Fatalf("body missing %s: %s", k, body)
		}
	}
}
