package site

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Sink struct {
	name       string
	updates    string
	daily      string
	maxUpdates int
	maxDays    int
	loc        *time.Location
}

type updatesFile struct {
	GeneratedAt string         `json:"generated_at"`
	Count       int            `json:"count"`
	Updates     []updateRecord `json:"updates"`
}

type updateRecord struct {
	ID        string   `json:"id"`
	RunAt     string   `json:"run_at"`
	FeedTitle string   `json:"feed_title"`
	Title     string   `json:"title"`
	Link      string   `json:"link"`
	Published string   `json:"published"`
	Tags      []string `json:"tags,omitempty"`
	TLDR      string   `json:"tldr,omitempty"`
	What      string   `json:"what,omitempty"`
	Why       string   `json:"why,omitempty"`
	Action    string   `json:"action,omitempty"`
	Tweet     string   `json:"tweet,omitempty"`
}

type dailyFile struct {
	GeneratedAt string         `json:"generated_at"`
	Days        []dailySummary `json:"days"`
}

type dailySummary struct {
	Date        string   `json:"date"`
	Total       int      `json:"total"`
	Feeds       []string `json:"feeds"`
	Highlights  []string `json:"highlights"`
	KeyPoints   []string `json:"key_points"`
	TweetDrafts []string `json:"tweet_drafts"`
}

func New(name string, cfg map[string]string) (*Sink, error) {
	updates := strings.TrimSpace(cfg["updates_path"])
	if updates == "" {
		updates = "site/data/updates.json"
	}
	daily := strings.TrimSpace(cfg["daily_path"])
	if daily == "" {
		daily = "site/data/daily.json"
	}
	maxUpdates := intFrom(cfg["max_updates"], 1200)
	maxDays := intFrom(cfg["max_days"], 30)

	tz := strings.TrimSpace(cfg["timezone"])
	if tz == "" {
		tz = "Europe/London"
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc = time.UTC
	}

	return &Sink{
		name:       name,
		updates:    updates,
		daily:      daily,
		maxUpdates: maxUpdates,
		maxDays:    maxDays,
		loc:        loc,
	}, nil
}

func (s *Sink) Name() string { return s.name }

func (s *Sink) Send(_ context.Context, batch core.Batch) error {
	current, err := s.loadUpdates()
	if err != nil {
		return err
	}

	byID := map[string]updateRecord{}
	for _, u := range current.Updates {
		byID[u.ID] = u
	}

	for _, e := range batch.Entries {
		rec := toRecord(batch.RunAt, e)
		byID[rec.ID] = rec
	}

	merged := make([]updateRecord, 0, len(byID))
	for _, u := range byID {
		merged = append(merged, u)
	}
	sort.Slice(merged, func(i, j int) bool {
		return timeKey(merged[i]).After(timeKey(merged[j]))
	})
	if len(merged) > s.maxUpdates {
		merged = merged[:s.maxUpdates]
	}

	out := updatesFile{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Count:       len(merged),
		Updates:     merged,
	}
	if err := writeJSON(s.updates, out); err != nil {
		return err
	}

	daily := s.buildDaily(merged)
	if err := writeJSON(s.daily, daily); err != nil {
		return err
	}
	return nil
}

func (s *Sink) loadUpdates() (updatesFile, error) {
	var out updatesFile
	b, err := os.ReadFile(s.updates)
	if os.IsNotExist(err) {
		return out, nil
	}
	if err != nil {
		return out, err
	}
	if len(b) == 0 {
		return out, nil
	}
	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}
	return out, nil
}

func (s *Sink) buildDaily(all []updateRecord) dailyFile {
	type bucket struct {
		feeds map[string]int
		items []updateRecord
	}
	byDay := map[string]*bucket{}
	for _, u := range all {
		t := timeKey(u).In(s.loc)
		day := t.Format("2006-01-02")
		b := byDay[day]
		if b == nil {
			b = &bucket{feeds: map[string]int{}, items: make([]updateRecord, 0, 16)}
			byDay[day] = b
		}
		b.items = append(b.items, u)
		if u.FeedTitle != "" {
			b.feeds[u.FeedTitle]++
		}
	}

	days := make([]string, 0, len(byDay))
	for d := range byDay {
		days = append(days, d)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(days)))
	if len(days) > s.maxDays {
		days = days[:s.maxDays]
	}

	result := make([]dailySummary, 0, len(days))
	for _, day := range days {
		b := byDay[day]
		sort.Slice(b.items, func(i, j int) bool {
			return timeKey(b.items[i]).After(timeKey(b.items[j]))
		})

		feeds := topFeedNames(b.feeds, 6)
		highlights := make([]string, 0, min(6, len(b.items)))
		keyPoints := make([]string, 0, 4)
		tweets := make([]string, 0, 3)

		for _, it := range b.items {
			if len(highlights) < 6 {
				highlights = append(highlights, fmt.Sprintf("%s · %s", it.FeedTitle, it.Title))
			}
			if it.TLDR != "" && len(keyPoints) < 4 {
				keyPoints = append(keyPoints, it.TLDR)
			}
			if it.Tweet != "" && len(tweets) < 3 {
				tweets = append(tweets, it.Tweet)
			}
		}

		result = append(result, dailySummary{
			Date:        day,
			Total:       len(b.items),
			Feeds:       feeds,
			Highlights:  highlights,
			KeyPoints:   keyPoints,
			TweetDrafts: tweets,
		})
	}

	return dailyFile{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Days:        result,
	}
}

func toRecord(runAt time.Time, e core.Entry) updateRecord {
	id := strings.TrimSpace(e.Link)
	if id == "" {
		id = strings.TrimSpace(e.Title) + "|" + e.Published.UTC().Format(time.RFC3339)
	}
	u := updateRecord{
		ID:        strings.ToLower(id),
		RunAt:     runAt.UTC().Format(time.RFC3339),
		FeedTitle: e.FeedTitle,
		Title:     e.Title,
		Link:      e.Link,
		Published: e.Published.UTC().Format(time.RFC3339),
		Tags:      e.Tags,
	}
	if e.Analysis != nil {
		u.TLDR = e.Analysis.TLDR
		u.What = e.Analysis.What
		u.Why = e.Analysis.Why
		u.Action = e.Analysis.Action
		u.Tweet = e.Analysis.Tweet
	}
	return u
}

func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func intFrom(v string, def int) int {
	if strings.TrimSpace(v) == "" {
		return def
	}
	n, err := strconv.Atoi(strings.TrimSpace(v))
	if err != nil || n <= 0 {
		return def
	}
	return n
}

func timeKey(u updateRecord) time.Time {
	if t, err := time.Parse(time.RFC3339, u.Published); err == nil {
		return t
	}
	if t, err := time.Parse(time.RFC3339, u.RunAt); err == nil {
		return t
	}
	return time.Unix(0, 0)
}

func topFeedNames(m map[string]int, n int) []string {
	type kv struct {
		name  string
		count int
	}
	arr := make([]kv, 0, len(m))
	for k, v := range m {
		arr = append(arr, kv{name: k, count: v})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].count == arr[j].count {
			return arr[i].name < arr[j].name
		}
		return arr[i].count > arr[j].count
	})
	if len(arr) > n {
		arr = arr[:n]
	}
	out := make([]string, 0, len(arr))
	for _, item := range arr {
		out = append(out, fmt.Sprintf("%s (%d)", item.name, item.count))
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
