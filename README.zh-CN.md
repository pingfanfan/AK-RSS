# opmlwatch

把 OPML 订阅清单变成可执行通知：规则过滤 + 每日摘要 + 可选 AI 简要分析。

## 项目做什么

`opmlwatch` 专注在“信息流 -> 行动”这条链路：

- 从 OPML 导入订阅
- 抓取每个订阅的最新条目（不是只看 feed 元信息）
- 用规则筛选（关键词/域名/内容长度）
- 可选 AI 分析（中英双语 `TLDR / WHAT / WHY / ACTION` + 平台发布文案）
- 输出到 `webhook`、`markdown`、`email`
- 用持久化已读状态去重（`.opmlwatch/seen_links.json`）

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

命中条目后，邮件会包含 AI 的 `TLDR/WHAT/WHY/ACTION/TWEET` 分析。

## 每 30 分钟自动执行（GitHub Action）

内置工作流：`.github/workflows/digest.yml`

- 每 30 分钟执行一次（`*/30 * * * *`）
- 抓取最新博客条目
- 生成可读摘要 + 可直接发社交媒体的 `TWEET` 草稿
- 每轮有处理上限（`max_entries_per_run`），避免首次跑爆量
- 有新命中条目时自动发邮件提醒

## 每 2 小时订阅邮件

工作流：`.github/workflows/subscriber-email.yml`

- 每 2 小时执行一次（`0 */2 * * *`）
- 从“订阅 issue”中读取订阅邮箱
- 仅发送“上次订阅推送之后”的新更新
- 用户关闭自己的订阅 issue 即取消订阅

## GitHub Pages 可视化页面

- 静态页面源码：`site/`
- 流程会自动更新数据文件：
  - `site/data/updates.json`（滚动更新）
  - `site/data/daily.json`（每日总结）
- 页面内置中英文一键切换（保留原文标题，并在切换语言时展示翻译）
- 每日自动生成可直接复制的社交平台文案（`X` / `LinkedIn` / `Threads`，中英双语，含来源链接）
- 页面发布工作流：`.github/workflows/pages.yml`

## 配置与接口

- 配置 schema：`config/schema.json`
- Sink 接口：`internal/sinks/sink.go`
- Summarizer 接口：`internal/summarizers/summarizer.go`
- 示例配置：`examples/config.basic.yaml`、`examples/config.digest.yaml`

## GitHub Actions

示例工作流：`examples/github-action/opmlwatch.yml`

## 许可证

Apache-2.0
