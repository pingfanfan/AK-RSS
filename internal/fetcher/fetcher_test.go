package fetcher

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opmlwatch/opmlwatch/internal/opml"
)

func TestFetchRSS(t *testing.T) {
	rss := `<?xml version="1.0"?><rss version="2.0"><channel><title>Test Feed</title><item><title>A</title><link>https://example.com/a</link><description>Body A</description></item></channel></rss>`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, rss)
	}))
	defer ts.Close()

	f := New(5)
	entries, err := f.Fetch(context.Background(), []opml.Feed{{Title: "Feed", XMLURL: ts.URL}}, Options{MaxItemsPerFeed: 1})
	if err != nil {
		t.Fatalf("fetch: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Title != "A" || entries[0].Link != "https://example.com/a" {
		t.Fatalf("unexpected entry: %+v", entries[0])
	}
}
