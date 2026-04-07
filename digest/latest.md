# OPMLWatch Digest

Generated at: 2026-04-07 10:46:01 UTC

- [Who Built This?](https://nesbitt.io/2026/04/07/who-built-this.html) (`nesbitt.io`)
  - TLDR: Go默认嵌入VCS信息，其他语言需手动配置。版本追踪对生产调试至关重要。
  - WHAT: 文章对比Go、Rust、.NET在二进制嵌入版本信息方面的做法。Go开箱即用，Rust依赖build.rs，.NET通过SourceLink接近Go。
  - WHY: 快速定位生产问题，避免猜测版本；增强安全审计，追踪代码来源；符合基础设施可观测性趋势。
  - ACTION: 评估项目：若用Go，直接使用`go version -m`；若用Rust/.NET，集成vergen或SourceLink；统一团队标准，确保所有服务可查询版本。
  - TWEET: Go默认嵌入版本信息，生产调试神器！对比Rust/.NET的手动配置，你的项目做到了吗？#开发实践 #DevOps
