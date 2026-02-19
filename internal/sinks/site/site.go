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
	"unicode/utf8"

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

	TitleZh string `json:"title_zh,omitempty"`
	TitleEn string `json:"title_en,omitempty"`

	TLDR   string `json:"tldr,omitempty"`
	What   string `json:"what,omitempty"`
	Why    string `json:"why,omitempty"`
	Action string `json:"action,omitempty"`
	Tweet  string `json:"tweet,omitempty"`

	TLDRZh   string `json:"tldr_zh,omitempty"`
	TLDREn   string `json:"tldr_en,omitempty"`
	WhatZh   string `json:"what_zh,omitempty"`
	WhatEn   string `json:"what_en,omitempty"`
	WhyZh    string `json:"why_zh,omitempty"`
	WhyEn    string `json:"why_en,omitempty"`
	ActionZh string `json:"action_zh,omitempty"`
	ActionEn string `json:"action_en,omitempty"`
	TweetZh  string `json:"tweet_zh,omitempty"`
	TweetEn  string `json:"tweet_en,omitempty"`

	XZh        string `json:"x_zh,omitempty"`
	XEn        string `json:"x_en,omitempty"`
	LinkedInZh string `json:"linkedin_zh,omitempty"`
	LinkedInEn string `json:"linkedin_en,omitempty"`
	ThreadsZh  string `json:"threads_zh,omitempty"`
	ThreadsEn  string `json:"threads_en,omitempty"`
}

type dailyFile struct {
	GeneratedAt string         `json:"generated_at"`
	Days        []dailySummary `json:"days"`
}

type dailySummary struct {
	Date        string           `json:"date"`
	Total       int              `json:"total"`
	Feeds       []string         `json:"feeds"`
	Highlights  []string         `json:"highlights"`
	KeyPoints   []string         `json:"key_points"`
	TweetDrafts []string         `json:"tweet_drafts"`
	SummaryZh   string           `json:"summary_zh,omitempty"`
	SummaryEn   string           `json:"summary_en,omitempty"`
	SocialPosts dailySocialPosts `json:"social_posts"`
}

type localizedPost struct {
	Zh string `json:"zh,omitempty"`
	En string `json:"en,omitempty"`
}

type dailySocialPosts struct {
	X        localizedPost `json:"x"`
	LinkedIn localizedPost `json:"linkedin"`
	Threads  localizedPost `json:"threads"`
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
	for i := range merged {
		merged[i] = normalizeRecord(merged[i])
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
		keyPointsZh := make([]string, 0, 4)
		keyPointsEn := make([]string, 0, 4)
		tweets := make([]string, 0, 3)

		for _, it := range b.items {
			if len(highlights) < 6 {
				highlights = append(highlights, fmt.Sprintf("%s · %s", it.FeedTitle, it.Title))
			}
			appendUniqueLimited(&keyPoints, firstNonEmpty(it.TLDRZh, it.TLDR, it.WhatZh, it.What), 4)
			appendUniqueLimited(&keyPointsZh, firstNonEmpty(it.TLDRZh, it.TLDR, it.WhatZh, it.What, it.TitleZh, it.Title), 4)
			appendUniqueLimited(&keyPointsEn, firstNonEmpty(it.TLDREn, it.WhatEn, it.TitleEn, it.Title, it.TLDR, it.What), 4)
			appendUniqueLimited(&tweets, firstNonEmpty(it.XZh, it.TweetZh, it.Tweet, it.XEn, it.TweetEn), 3)
		}

		summaryZh := buildDailySummaryText(day, keyPointsZh, "zh")
		summaryEn := buildDailySummaryText(day, keyPointsEn, "en")
		social := buildDailySocialPosts(day, keyPointsZh, keyPointsEn, highlights)

		result = append(result, dailySummary{
			Date:        day,
			Total:       len(b.items),
			Feeds:       feeds,
			Highlights:  highlights,
			KeyPoints:   keyPoints,
			TweetDrafts: tweets,
			SummaryZh:   summaryZh,
			SummaryEn:   summaryEn,
			SocialPosts: social,
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
		u.TitleZh = firstNonEmpty(e.Analysis.TitleZh, e.Title)
		u.TitleEn = firstNonEmpty(e.Analysis.TitleEn, e.Title)

		u.TLDR = firstNonEmpty(e.Analysis.TLDR, e.Analysis.TLDRZh, e.Analysis.TLDREn)
		u.What = firstNonEmpty(e.Analysis.What, e.Analysis.WhatZh, e.Analysis.WhatEn)
		u.Why = firstNonEmpty(e.Analysis.Why, e.Analysis.WhyZh, e.Analysis.WhyEn)
		u.Action = firstNonEmpty(e.Analysis.Action, e.Analysis.ActionZh, e.Analysis.ActionEn)
		u.Tweet = firstNonEmpty(e.Analysis.Tweet, e.Analysis.XZh, e.Analysis.TweetZh, e.Analysis.XEn, e.Analysis.TweetEn)

		u.TLDRZh = firstNonEmpty(e.Analysis.TLDRZh, e.Analysis.TLDR)
		u.TLDREn = firstNonEmpty(e.Analysis.TLDREn, e.Analysis.TLDR)
		u.WhatZh = firstNonEmpty(e.Analysis.WhatZh, e.Analysis.What)
		u.WhatEn = firstNonEmpty(e.Analysis.WhatEn, e.Analysis.What)
		u.WhyZh = firstNonEmpty(e.Analysis.WhyZh, e.Analysis.Why)
		u.WhyEn = firstNonEmpty(e.Analysis.WhyEn, e.Analysis.Why)
		u.ActionZh = firstNonEmpty(e.Analysis.ActionZh, e.Analysis.Action)
		u.ActionEn = firstNonEmpty(e.Analysis.ActionEn, e.Analysis.Action)
		u.TweetZh = firstNonEmpty(e.Analysis.TweetZh, e.Analysis.Tweet, e.Analysis.XZh)
		u.TweetEn = firstNonEmpty(e.Analysis.TweetEn, e.Analysis.Tweet, e.Analysis.XEn)

		u.XZh = firstNonEmpty(e.Analysis.XZh, e.Analysis.TweetZh, e.Analysis.Tweet)
		u.XEn = firstNonEmpty(e.Analysis.XEn, e.Analysis.TweetEn, e.Analysis.Tweet)
		u.LinkedInZh = firstNonEmpty(e.Analysis.LinkedInZh, u.XZh)
		u.LinkedInEn = firstNonEmpty(e.Analysis.LinkedInEn, u.XEn)
		u.ThreadsZh = firstNonEmpty(e.Analysis.ThreadsZh, u.XZh)
		u.ThreadsEn = firstNonEmpty(e.Analysis.ThreadsEn, u.XEn)
	} else {
		u.TitleZh = e.Title
		u.TitleEn = e.Title
	}
	return normalizeRecord(u)
}

func normalizeRecord(u updateRecord) updateRecord {
	u.TitleZh = firstNonEmpty(u.TitleZh, u.Title)
	u.TitleEn = firstNonEmpty(u.TitleEn, u.Title)

	u.TLDRZh = firstNonEmpty(u.TLDRZh, u.TLDR, u.WhatZh, u.What, u.TitleZh, u.Title)
	u.TLDREn = firstNonEmpty(u.TLDREn, u.WhatEn, u.TitleEn, u.Title, u.TLDR, u.What)
	u.WhatZh = firstNonEmpty(u.WhatZh, u.What, u.TLDRZh)
	u.WhatEn = firstNonEmpty(u.WhatEn, u.TitleEn, u.What, u.TLDREn)
	u.WhyZh = firstNonEmpty(u.WhyZh, u.Why, "这条更新可能影响你的技术判断。")
	u.WhyEn = firstNonEmpty(u.WhyEn, u.Why, "This update can affect your technical decisions.")
	u.ActionZh = firstNonEmpty(u.ActionZh, u.Action, "打开原文并判断是否要进入你的行动清单。")
	u.ActionEn = firstNonEmpty(u.ActionEn, u.Action, "Open the source and decide whether to add it to your action list.")

	fallbackPost := firstNonEmpty(u.Title, u.Link)
	u.TweetZh = firstNonEmpty(u.TweetZh, u.XZh, u.Tweet, fallbackPost)
	u.TweetEn = firstNonEmpty(u.TweetEn, u.XEn, u.Tweet, fallbackPost)
	u.XZh = firstNonEmpty(u.XZh, u.TweetZh, u.Tweet, fallbackPost)
	u.XEn = firstNonEmpty(u.XEn, u.TweetEn, fallbackPost, u.Tweet)
	u.LinkedInZh = firstNonEmpty(u.LinkedInZh, u.XZh, u.TweetZh)
	u.LinkedInEn = firstNonEmpty(u.LinkedInEn, u.XEn, u.TweetEn)
	u.ThreadsZh = firstNonEmpty(u.ThreadsZh, u.XZh, u.TweetZh)
	u.ThreadsEn = firstNonEmpty(u.ThreadsEn, u.XEn, u.TweetEn)

	u.TLDR = firstNonEmpty(u.TLDR, u.TLDRZh, u.TLDREn)
	u.What = firstNonEmpty(u.What, u.WhatZh, u.WhatEn)
	u.Why = firstNonEmpty(u.Why, u.WhyZh, u.WhyEn)
	u.Action = firstNonEmpty(u.Action, u.ActionZh, u.ActionEn)
	u.Tweet = firstNonEmpty(u.Tweet, u.XZh, u.XEn, u.TweetZh, u.TweetEn)
	return u
}

func buildDailySummaryText(day string, points []string, lang string) string {
	use := make([]string, 0, 3)
	for _, p := range points {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		use = append(use, p)
		if len(use) == 3 {
			break
		}
	}
	if len(use) == 0 {
		return ""
	}
	if lang == "zh" {
		return clipRunes(fmt.Sprintf("%s 技术日报：%s。", day, strings.Join(use, "；")), 320)
	}
	return clipRunes(fmt.Sprintf("Daily brief (%s): %s.", day, strings.Join(use, "; ")), 320)
}

func buildDailySocialPosts(day string, zhPoints, enPoints, highlights []string) dailySocialPosts {
	zh := firstThree(zhPoints)
	en := firstThree(enPoints)
	if len(zh) == 0 {
		zh = firstThree(highlights)
	}
	if len(en) == 0 {
		en = firstThree(highlights)
	}

	zhJoined := strings.Join(zh, "；")
	enJoined := strings.Join(en, "; ")

	zhBullets := bullets(zh)
	enBullets := bullets(en)

	return dailySocialPosts{
		X: localizedPost{
			Zh: clipRunes(fmt.Sprintf("【%s 日报】%s。我把可执行建议整理好了，欢迎转发讨论。#技术趋势 #AI #工程实践", day, zhJoined), 260),
			En: clipRunes(fmt.Sprintf("AK-RSS Daily (%s): %s. Actionable takeaways included. What would you ship next? #AI #Engineering #BuildInPublic", day, enJoined), 260),
		},
		LinkedIn: localizedPost{
			Zh: strings.TrimSpace(fmt.Sprintf(
				"今天的技术情报（%s）\n%s\n\n我整理了可执行动作与原文链接，适合团队同步或选题参考。\n#AI #工程实践 #产品洞察",
				day, zhBullets,
			)),
			En: strings.TrimSpace(fmt.Sprintf(
				"Tech signal brief (%s)\n%s\n\nI compiled actionable next steps and source links for team sync and content planning.\n#AI #Engineering #Product",
				day, enBullets,
			)),
		},
		Threads: localizedPost{
			Zh: clipRunes(fmt.Sprintf(
				"今天最值得看的 3 个点：%s。\n我把行动建议做成日报了，可以直接拿去用。",
				zhJoined,
			), 320),
			En: clipRunes(fmt.Sprintf(
				"Three signals worth your attention today: %s.\nI turned them into an action-ready daily brief you can use right away.",
				enJoined,
			), 320),
		},
	}
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

func firstThree(items []string) []string {
	out := make([]string, 0, 3)
	for _, it := range items {
		it = strings.TrimSpace(it)
		if it == "" {
			continue
		}
		out = append(out, it)
		if len(out) == 3 {
			break
		}
	}
	return out
}

func bullets(items []string) string {
	p := firstThree(items)
	if len(p) == 0 {
		return "- No major updates captured"
	}
	lines := make([]string, 0, len(p))
	for _, it := range p {
		lines = append(lines, "- "+it)
	}
	return strings.Join(lines, "\n")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		v = strings.TrimSpace(v)
		if v != "" {
			return v
		}
	}
	return ""
}

func appendUniqueLimited(dst *[]string, value string, limit int) {
	value = strings.TrimSpace(value)
	if value == "" || len(*dst) >= limit {
		return
	}
	for _, cur := range *dst {
		if cur == value {
			return
		}
	}
	*dst = append(*dst, value)
}

func clipRunes(s string, max int) string {
	if max <= 0 {
		return ""
	}
	s = strings.TrimSpace(s)
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	r := []rune(s)
	if len(r) <= max {
		return string(r)
	}
	return strings.TrimSpace(string(r[:max]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
