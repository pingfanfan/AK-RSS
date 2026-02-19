package core

import "time"

type Entry struct {
	FeedTitle string    `json:"feed_title"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Author    string    `json:"author,omitempty"`
	Published time.Time `json:"published"`
	Content   string    `json:"content,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Analysis  *Analysis `json:"analysis,omitempty"`
}

type Batch struct {
	RunAt   time.Time         `json:"run_at"`
	Entries []Entry           `json:"entries"`
	Meta    map[string]string `json:"meta,omitempty"`
}

type Analysis struct {
	TLDR   string `json:"tldr,omitempty"`
	What   string `json:"what"`
	Why    string `json:"why"`
	Action string `json:"action"`
	Tweet  string `json:"tweet,omitempty"`

	TitleZh  string `json:"title_zh,omitempty"`
	TitleEn  string `json:"title_en,omitempty"`
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

	Model string `json:"model,omitempty"`
}
