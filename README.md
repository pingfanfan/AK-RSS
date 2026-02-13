# opmlwatch

Turn OPML into actionable updates: rule-based alerts + daily digests + optional AI summaries.

把 OPML 订阅清单变成可执行通知：规则过滤 + 每日摘要 + 可选 AI 简要分析。

## Docs

- English: `README.en.md`
- 中文: `README.zh-CN.md`

## Quick Start

```bash
docker compose build
docker compose run --rm opmlwatch
```

Output: `digest/latest.md`

## Repository

- Config schema: `config/schema.json`
- Example config: `examples/config.digest.yaml`
- GitHub Action template: `examples/github-action/opmlwatch.yml`

## License

Apache-2.0
