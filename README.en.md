# opmlwatch

Turn OPML into actionable updates: rule-based alerts + daily digests + optional AI summaries.

## What It Does

`opmlwatch` focuses on the "feed -> action" pipeline:

- Import feeds from OPML
- Fetch latest entries from each feed (not just feed metadata)
- Apply rules (keywords/domains/content length)
- Optional AI analysis (`TLDR / WHAT / WHY / ACTION / TWEET`)
- Deliver to sinks (`webhook`, `markdown`, `email`)
- Deduplicate with persisted seen-state (`.opmlwatch/seen_links.json`)

## Quickstart

### Option A: Docker (No local Go)

```bash
docker compose build
docker compose run --rm opmlwatch
```

Output: `digest/latest.md`

### Option B: Local Go

```bash
go build ./cmd/opmlwatch
./opmlwatch validate-config --config examples/config.digest.yaml
./opmlwatch run --config examples/config.digest.yaml --dry-run
```

## Email Notification + AI Summary

Edit `examples/config.digest.yaml`:

1. Enable summarizer:

```yaml
summarizer:
  type: "openai"
  enabled: true
  config:
    base_url: "https://api.openai.com/v1"
    model: "gpt-4o-mini"
    api_key: "env:OPENAI_API_KEY"
```

2. Enable email sink:

```yaml
- type: "email"
  name: "mail-alert"
  enabled: true
  config:
    host: "smtp.example.com"
    port: "587"
    username: "your-user"
    password: "your-password"
    from: "bot@example.com"
    to: "you@example.com"
```

Run:

```bash
export OPENAI_API_KEY="your_key"
docker compose run --rm opmlwatch
```

When entries match, email includes AI `TLDR/WHAT/WHY/ACTION/TWEET` lines.

## 30-Minute GitHub Action

Built-in workflow: `.github/workflows/digest.yml`

- Runs every 30 minutes (`*/30 * * * *`)
- Pulls latest feed entries
- Generates digest + tweet-ready summary lines
- Applies per-run cap (`max_entries_per_run`) to prevent first-run overload
- Sends email alert when there are new matched entries

## 2-Hour Subscriber Emails

Workflow: `.github/workflows/subscriber-email.yml`

- Runs every 2 hours (`0 */2 * * *`)
- Reads subscriber emails from open subscription issues
- Sends only updates newer than the last subscriber-send cursor
- Unsubscribe by closing your subscription issue

## GitHub Pages Dashboard

- Static site source: `site/`
- Data files updated by the pipeline:
  - `site/data/updates.json` (rolling updates)
  - `site/data/daily.json` (daily summaries)
- Deployment workflow: `.github/workflows/pages.yml`

## Config and Interfaces

- Config schema: `config/schema.json`
- Sink interface: `internal/sinks/sink.go`
- Summarizer interface: `internal/summarizers/summarizer.go`
- Example configs: `examples/config.basic.yaml`, `examples/config.digest.yaml`

## GitHub Actions

Starter workflow: `examples/github-action/opmlwatch.yml`

## License

Apache-2.0
