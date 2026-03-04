# OPMLWatch Digest

Generated at: 2026-03-04 10:43:08 UTC

- [Package Managers Need to Cool Down](https://nesbitt.io/2026/03/04/package-managers-need-to-cool-down.html) (`nesbitt.io`)
  - TLDR: 依赖冷却（如设置7天等待期）能有效阻断自动化供应链攻击。主流包管理器（pnpm、Yarn、Bun、Go、Rust、Python）已支持该功能，开发者应立即在项目中配置。
  - WHAT: 依赖冷却是一种安全策略，要求包管理器在安装新发布的包版本前，强制等待一个预设的最小时间（如7天）。这为安全社区和工具留出检测恶意版本的时间窗口，防止自动化流程在攻击者发布恶意包后立即将其拉取到数千个项目中。不同工具的实现名称各异（如pnpm的minimumReleaseAge、Yarn的npmMinimalAgeGate），但核心逻辑一致。
  - WHY: 攻击者常通过窃取维护者凭证或接管休眠包，发布恶意版本后利用CI/CD等自动化工具在数小时内扩散。研究显示，多数供应链攻击的“黄金窗口”短于一周。依赖冷却通过人为引入延迟，将响应时间拉回人类可干预的范畴，是成本极低且效果显著的安全加固手段。
  - ACTION: 1. 检查项目所用包管理器版本是否支持冷却功能（如pnpm≥10.16、Yarn≥4.10.0、Bun≥1.2、Go≥1.23、Rust/Cargo、Python/pip）。2. 在配置文件（如.pnpmfile.cjs、.yarnrc.yml、.bunfig.toml、go.mod、Cargo.toml、requirements.txt）中启用冷却设置，建议初始值为7天。3. 为高度信任的核心依赖配置排除列表（如pnpm的minimumReleaseAgeExclude、Yarn的npmPreapprovedPackages），避免影响关键安全更新。4. 在CI/CD流水线中验证设置生效，并监控因冷却导致的安装延迟。
  - TWEET: 你的CI/CD是否在自动引入风险？依赖冷却（Dependency Cooldown）是防御供应链攻击的隐藏王牌。通过强制等待新包版本（如7天），为安全响应争取时间。主流工具已支持，配置指南见原文。安全左移，从依赖管理开始。
