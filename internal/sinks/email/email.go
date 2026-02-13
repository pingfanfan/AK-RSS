package email

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/core"
)

type Sink struct {
	name          string
	host          string
	port          int
	username      string
	password      string
	from          string
	to            []string
	subjectPrefix string
}

func New(name string, cfg map[string]string) (*Sink, error) {
	port := 587
	if p := resolve(cfg["port"]); p != "" {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("email sink %s invalid port: %w", name, err)
		}
		port = n
	}
	to := splitCSV(resolve(cfg["to"]))
	s := &Sink{
		name:          name,
		host:          resolve(cfg["host"]),
		port:          port,
		username:      resolve(cfg["username"]),
		password:      resolve(cfg["password"]),
		from:          resolve(cfg["from"]),
		to:            to,
		subjectPrefix: resolve(cfg["subject_prefix"]),
	}
	if s.host == "" || s.from == "" || len(s.to) == 0 {
		return nil, fmt.Errorf("email sink %s requires host/from/to", name)
	}
	return s, nil
}

func (s *Sink) Name() string { return s.name }

func (s *Sink) Send(_ context.Context, batch core.Batch) error {
	if len(batch.Entries) == 0 {
		return nil
	}
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	subject := "OPMLWatch updates"
	if s.subjectPrefix != "" {
		subject = s.subjectPrefix + " " + subject
	}

	body := buildBody(batch)
	msg := "From: " + s.from + "\r\n" +
		"To: " + strings.Join(s.to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" + body

	var auth smtp.Auth
	if s.username != "" && s.password != "" {
		auth = smtp.PlainAuth("", s.username, s.password, s.host)
	}
	return smtp.SendMail(addr, auth, s.from, s.to, []byte(msg))
}

func buildBody(batch core.Batch) string {
	var b strings.Builder
	b.WriteString("OPMLWatch\n")
	b.WriteString("Generated at: " + batch.RunAt.Format(time.RFC3339) + "\n\n")
	for _, e := range batch.Entries {
		b.WriteString("- " + e.Title + "\n")
		b.WriteString("  Link: " + e.Link + "\n")
		if e.Analysis != nil {
			if e.Analysis.TLDR != "" {
				b.WriteString("  TLDR: " + e.Analysis.TLDR + "\n")
			}
			b.WriteString("  WHAT: " + e.Analysis.What + "\n")
			b.WriteString("  WHY: " + e.Analysis.Why + "\n")
			b.WriteString("  ACTION: " + e.Analysis.Action + "\n")
			if e.Analysis.Tweet != "" {
				b.WriteString("  TWEET: " + e.Analysis.Tweet + "\n")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func splitCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func resolve(v string) string {
	v = strings.TrimSpace(v)
	if strings.HasPrefix(v, "env:") {
		return strings.TrimSpace(os.Getenv(strings.TrimPrefix(v, "env:")))
	}
	return v
}
