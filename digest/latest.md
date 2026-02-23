# OPMLWatch Digest

Generated at: 2026-02-23 10:45:05 UTC

- [Where Do Specifications Fit in the Dependency Tree?](https://nesbitt.io/2026/02/23/where-do-specifications-fit-in-the-dependency-tree.html) (`nesbitt.io`)
  - TLDR: 语言规范（如Ruby 3.0）未被纳入依赖管理，运行时版本声明碎片化且 enforcement 不一致，导致依赖重复声明与冲突。
  - WHAT: 当前依赖管理（manifest/lockfile/SBOM）只追踪库依赖，忽略语言/协议规范；运行时版本通过 gemspec/Gemfile/.tool-versions 等多处声明，机制各异且可能冲突。
  - WHY: 导致环境不一致、安全漏洞难追踪（如JRuby谎报版本）、开发者认知负荷增加，破坏可重复构建与供应链安全。
  - ACTION: 1. 在项目中统一运行时声明（如仅用 .tool-versions）；2. 推动工具链支持规范追踪（如SBOM扩展）；3. 审计现有项目，消除冲突声明。
  - TWEET: 你的项目有多个地方声明Ruby版本吗？gemspec、Gemfile、.tool-versions 可能冲突！语言规范（RFC/ECMA）甚至不在SBOM里。如何解决？讨论→
