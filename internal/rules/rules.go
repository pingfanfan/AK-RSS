package rules

import (
	"net/url"
	"strings"

	"github.com/opmlwatch/opmlwatch/internal/config"
	"github.com/opmlwatch/opmlwatch/internal/core"
)

func Apply(entries []core.Entry, rules []config.RuleConfig) []core.Entry {
	if len(rules) == 0 {
		return entries
	}
	out := make([]core.Entry, 0, len(entries))
	for _, entry := range entries {
		if matchAny(entry, rules) {
			out = append(out, entry)
		}
	}
	return out
}

func matchAny(entry core.Entry, rules []config.RuleConfig) bool {
	for _, rule := range rules {
		if matchRule(entry, rule) {
			return true
		}
	}
	return false
}

func matchRule(entry core.Entry, rule config.RuleConfig) bool {
	text := strings.ToLower(entry.Title + "\n" + entry.Content)
	for _, ex := range rule.ExcludeKeywords {
		if strings.Contains(text, strings.ToLower(ex)) {
			return false
		}
	}
	if len(rule.IncludeKeywords) > 0 {
		ok := false
		for _, in := range rule.IncludeKeywords {
			if strings.Contains(text, strings.ToLower(in)) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	if rule.MinContentLength > 0 && len(entry.Content) < rule.MinContentLength {
		return false
	}
	if len(rule.IncludeDomains) > 0 {
		u, err := url.Parse(entry.Link)
		if err != nil {
			return false
		}
		host := strings.ToLower(u.Hostname())
		ok := false
		for _, d := range rule.IncludeDomains {
			d = strings.ToLower(strings.TrimSpace(d))
			if d != "" && (host == d || strings.HasSuffix(host, "."+d)) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
