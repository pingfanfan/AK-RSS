# OPMLWatch Digest

Generated at: 2026-03-10 10:43:52 UTC

- [Just Use Postgres](https://nesbitt.io/2026/03/10/just-use-postgres.html) (`nesbitt.io`)
  - TLDR: 通过 Postgres 扩展 omni_git 实现 git 智能 HTTP 协议，使数据库直接成为 git 主机、构建系统和运行时，无需反向代理、容器或独立应用服务器。
  - WHAT: 一个 Postgres 扩展（omni_git），在数据库内部实现 git 的智能 HTTP 协议，使 Postgres 能直接处理 git push/clone 操作，并结合 omnigres 将 SQL 文件自动部署为运行中的代码。
  - WHY: 消除反向代理、容器运行时等中间层，大幅简化架构和运维；利用 Postgres 的扩展性和可靠性，将版本控制、部署和运行时合一，降低复杂度和成本。
  - ACTION: 尝试 omni_git 扩展，体验 git push 直接部署；评估将 Postgres 扩展用于其他基础设施场景（如 API 网关、消息队列）。
  - TWEET: 刚看到一个疯狂的项目：用 Postgres 扩展让数据库直接支持 git push 部署！整个平台就一个进程，没有那些烦人的中间件。这可能是简化基础设施的未来方向。 #Postgres #Git
