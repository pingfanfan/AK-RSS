# Architecture

OPMLWatch is designed as a pipeline with stable extension points:

1. Fetcher: retrieve feed updates with cache headers (planned)
2. Normalizer: unify RSS/Atom/JSON Feed into `core.Entry` (planned)
3. Store: persist seen entries and run metadata (planned, SQLite in MVP)
4. Rules: filter and route entry flow
5. Sinks: send alerts, webhooks, or digest outputs
6. Summarizers: optional structured summary provider

Current MVP covers:
- OPML ingestion
- rule filtering
- sink delivery (`webhook`, `markdown`)
- configuration validation

## Extension Contracts

- Sink interface: `internal/sinks/sink.go`
- Summarizer interface: `internal/summarizers/summarizer.go`

These contracts are kept small on purpose to lower integration friction for contributors.
