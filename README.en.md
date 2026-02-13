# opmlwatch

Turn OPML into actionable updates: rule-based alerts + daily digests + optional AI summaries.

## What It Does

`opmlwatch` focuses on the "feed -> action" pipeline:

- Import feeds from OPML
- Apply rules (keywords/domains/content length)
- Optional AI analysis (`WHAT / WHY / ACTION`)
- Deliver to sinks (`webhook`, `markdown`, `email`)

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

When entries match, email includes AI `WHAT/WHY/ACTION` lines.

## Config and Interfaces

- Config schema: `config/schema.json`
- Sink interface: `internal/sinks/sink.go`
- Summarizer interface: `internal/summarizers/summarizer.go`
- Example configs: `examples/config.basic.yaml`, `examples/config.digest.yaml`

## GitHub Actions

Starter workflow: `examples/github-action/opmlwatch.yml`

## License

Apache-2.0
