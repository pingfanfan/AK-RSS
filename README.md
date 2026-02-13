# opmlwatch

Turn OPML into actionable updates: rule-based alerts + daily digests + optional AI summaries.

把 OPML 订阅清单变成可执行通知：规则过滤 + 每日摘要 + 可选 AI 简要分析。

## EN: What It Does

`opmlwatch` focuses on "feed -> action":

- Import feeds from OPML
- Apply rules (keywords/domains/content length)
- Optional AI analysis (`WHAT / WHY / ACTION`)
- Deliver to sinks (`webhook`, `markdown`, `email`)

## 中文：项目定位

`opmlwatch` 解决的是“信息流到可操作通知”这条链路：

- 从 OPML 读取订阅
- 用规则筛选真正重要的更新
- 可选调用 AI 生成简要分析（是什么 / 为什么重要 / 建议动作）
- 输出到工作流渠道（Webhook、Markdown 摘要、Email）

## EN: Quickstart

### Option A: Docker (No local Go)

```bash
docker compose build
docker compose run --rm opmlwatch
```

Output file: `digest/latest.md`

### Option B: Local Go

```bash
go build ./cmd/opmlwatch
./opmlwatch validate-config --config examples/config.digest.yaml
./opmlwatch run --config examples/config.digest.yaml --dry-run
```

## 中文：快速开始

### 方案 A：Docker（不需要本地 Go）

```bash
docker compose build
docker compose run --rm opmlwatch
```

运行后会生成：`digest/latest.md`

### 方案 B：本地 Go

```bash
go build ./cmd/opmlwatch
./opmlwatch validate-config --config examples/config.digest.yaml
./opmlwatch run --config examples/config.digest.yaml --dry-run
```

## EN: Email Notification + AI Analysis

Edit `examples/config.digest.yaml`:

1. Enable AI summarizer:

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

Then run:

```bash
export OPENAI_API_KEY="your_key"
docker compose run --rm opmlwatch
```

## 中文：邮件通知 + AI 简要分析

编辑 `examples/config.digest.yaml`：

1. 打开 AI 摘要（`summarizer.enabled: true`），并配置 API Key
2. 打开邮件 sink（`email.enabled: true`），填好 SMTP 参数

运行：

```bash
export OPENAI_API_KEY="你的密钥"
docker compose run --rm opmlwatch
```

有命中条目时，会发邮件并附带 AI 输出的 `WHAT / WHY / ACTION`。

## EN/中文: Config and Interfaces

- Config schema: `config/schema.json`
- Sink interface: `internal/sinks/sink.go`
- Summarizer interface: `internal/summarizers/summarizer.go`
- Example configs: `examples/config.basic.yaml`, `examples/config.digest.yaml`

## GitHub Action

Starter workflow: `examples/github-action/opmlwatch.yml`

## License

Apache-2.0
