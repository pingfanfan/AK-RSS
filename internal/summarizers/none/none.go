package none

import (
	"context"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Summarizer struct{}

func New() *Summarizer { return &Summarizer{} }

func (s *Summarizer) Name() string { return "none" }

func (s *Summarizer) Summarize(_ context.Context, entry core.Entry) (core.Analysis, error) {
	tweet := entry.Title + " " + entry.Link
	return core.Analysis{
		TLDR:   entry.Title,
		What:   entry.Title,
		Why:    "Fresh update from monitored feed.",
		Action: "Open source link and decide whether to read now or add to digest backlog.",
		Tweet:  tweet,

		TitleZh:  entry.Title,
		TitleEn:  entry.Title,
		TLDRZh:   entry.Title,
		TLDREn:   entry.Title,
		WhatZh:   entry.Title,
		WhatEn:   entry.Title,
		WhyZh:    "这是一条来自监控信源的新更新。",
		WhyEn:    "Fresh update from monitored feed.",
		ActionZh: "打开原文，决定立即阅读还是加入待办清单。",
		ActionEn: "Open source link and decide whether to read now or add to digest backlog.",
		TweetZh:  tweet,
		TweetEn:  tweet,

		XZh:        tweet,
		XEn:        tweet,
		LinkedInZh: tweet,
		LinkedInEn: tweet,
		ThreadsZh:  tweet,
		ThreadsEn:  tweet,
		Model:      "none",
	}, nil
}
