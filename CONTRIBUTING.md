# Contributing

Thanks for contributing to OPMLWatch.

## Local Development

- Go 1.22+
- Run `go build ./...`
- Run `go test ./...`
- Validate sample config: `go run ./cmd/opmlwatch validate-config --config examples/config.digest.yaml`

## Adding a New Sink

1. Add implementation under `internal/sinks/<name>/`
2. Implement `sinks.Sink`
3. Register constructor in `internal/sinks/sink.go`
4. Add config example in `examples/`

## Pull Requests

- Keep PRs focused and small
- Include tests for behavior changes
- Update README/docs for user-visible changes
