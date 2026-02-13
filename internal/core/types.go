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
	What   string `json:"what"`
	Why    string `json:"why"`
	Action string `json:"action"`
	Model  string `json:"model,omitempty"`
}
