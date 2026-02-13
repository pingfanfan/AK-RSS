package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/opmlwatch/opmlwatch/internal/core"
	"github.com/opmlwatch/opmlwatch/internal/opml"
)

type Options struct {
	MaxItemsPerFeed int
	TimeoutSeconds  int
	MaxConcurrent   int
}

type Fetcher struct {
	client *http.Client
}

func New(timeoutSeconds int) *Fetcher {
	if timeoutSeconds <= 0 {
		timeoutSeconds = 15
	}
	return &Fetcher{
		client: &http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second},
	}
}

func (f *Fetcher) Fetch(ctx context.Context, feeds []opml.Feed, opts Options) ([]core.Entry, error) {
	maxItems := opts.MaxItemsPerFeed
	if maxItems <= 0 {
		maxItems = 3
	}
	maxConcurrent := opts.MaxConcurrent
	if maxConcurrent <= 0 {
		maxConcurrent = 12
	}
	if maxConcurrent > len(feeds) {
		maxConcurrent = len(feeds)
	}
	if maxConcurrent == 0 {
		return nil, nil
	}

	entries := make([]core.Entry, 0, len(feeds)*maxItems)
	seen := map[string]struct{}{}
	var okCount int
	type result struct {
		entries []core.Entry
		err     error
	}
	jobs := make(chan opml.Feed)
	results := make(chan result, len(feeds))
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for feed := range jobs {
			batch, err := f.fetchOne(ctx, feed, maxItems)
			results <- result{entries: batch, err: err}
		}
	}
	for i := 0; i < maxConcurrent; i++ {
		wg.Add(1)
		go worker()
	}
	go func() {
		for _, feed := range feeds {
			jobs <- feed
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if r.err != nil {
			continue
		}
		okCount++
		for _, e := range r.entries {
			key := normalizeKey(e.Link, e.Title)
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = struct{}{}
			entries = append(entries, e)
		}
	}

	if okCount == 0 {
		return nil, fmt.Errorf("all feeds failed to fetch")
	}
	return entries, nil
}

func (f *Fetcher) fetchOne(ctx context.Context, feed opml.Feed, maxItems int) ([]core.Entry, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feed.XMLURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "opmlwatch/0.1 (+https://github.com/pingfanfan/AK-RSS)")
	req.Header.Set("Accept", "application/rss+xml, application/atom+xml, application/xml, text/xml;q=0.9, */*;q=0.8")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("feed %s status %s", feed.XMLURL, resp.Status)
	}

	data, err := io.ReadAll(io.LimitReader(resp.Body, 8*1024*1024))
	if err != nil {
		return nil, err
	}
	parsed, err := gofeed.NewParser().ParseString(string(data))
	if err != nil {
		return nil, err
	}

	feedTitle := strings.TrimSpace(feed.Title)
	if feedTitle == "" {
		feedTitle = strings.TrimSpace(parsed.Title)
	}

	out := make([]core.Entry, 0, maxItems)
	for _, item := range parsed.Items {
		if len(out) >= maxItems {
			break
		}
		link := strings.TrimSpace(item.Link)
		title := strings.TrimSpace(item.Title)
		if link == "" && title == "" {
			continue
		}

		published := time.Now().UTC()
		if item.PublishedParsed != nil {
			published = item.PublishedParsed.UTC()
		} else if item.UpdatedParsed != nil {
			published = item.UpdatedParsed.UTC()
		}

		author := ""
		if item.Author != nil {
			author = strings.TrimSpace(item.Author.Name)
		}

		content := firstNonEmpty(item.Content, item.Description)
		content = sanitizeText(content)

		tags := make([]string, 0, len(item.Categories))
		for _, c := range item.Categories {
			c = strings.TrimSpace(c)
			if c != "" {
				tags = append(tags, c)
			}
		}

		out = append(out, core.Entry{
			FeedTitle: feedTitle,
			Title:     title,
			Link:      link,
			Author:    author,
			Published: published,
			Content:   content,
			Tags:      tags,
		})
	}
	return out, nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func sanitizeText(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.Join(strings.Fields(s), " ")
	if len(s) > 4000 {
		return s[:4000]
	}
	return s
}

func normalizeKey(link, title string) string {
	if strings.TrimSpace(link) != "" {
		return strings.ToLower(strings.TrimSpace(link))
	}
	return strings.ToLower(strings.TrimSpace(title))
}
