# opmlwatch

把 OPML 订阅清单变成可执行通知：规则过滤 + 每日摘要 + 可选 AI 简要分析。

## 项目做什么

`opmlwatch` 专注在“信息流 -> 行动”这条链路：

- 从 OPML 导入订阅
- 用规则筛选（关键词/域名/内容长度）
- 可选 AI 分析（`WHAT / WHY / ACTION`）
- 输出到 `webhook`、`markdown`、`email`

## 快速开始

### 方案 A：Docker（不需要本地 Go）

```bash
docker compose build
docker compose run --rm opmlwatch
```

输出文件：`digest/latest.md`

### 方案 B：本地 Go

```bash
go build ./cmd/opmlwatch
./opmlwatch validate-config --config examples/config.digest.yaml
./opmlwatch run --config examples/config.digest.yaml --dry-run
```

## 邮件通知 + AI 简要分析

编辑 `examples/config.digest.yaml`：

1. 开启 summarizer：

```yaml
summarizer:
  type: "openai"
  enabled: true
  config:
    base_url: "https://api.openai.com/v1"
    model: "gpt-4o-mini"
    api_key: "env:OPENAI_API_KEY"
```

2. 开启 email sink：

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

运行：

```bash
export OPENAI_API_KEY="你的密钥"
docker compose run --rm opmlwatch
```

命中条目后，邮件会包含 AI 的 `WHAT/WHY/ACTION` 分析。

## 配置与接口

- 配置 schema：`config/schema.json`
- Sink 接口：`internal/sinks/sink.go`
- Summarizer 接口：`internal/summarizers/summarizer.go`
- 示例配置：`examples/config.basic.yaml`、`examples/config.digest.yaml`

## GitHub Actions

示例工作流：`examples/github-action/opmlwatch.yml`

## 许可证

Apache-2.0
