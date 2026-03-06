# OPMLWatch Digest

Generated at: 2026-03-06 20:42:19 UTC

- [Google’s Threat Intelligence Group on Coruna a Powerful iOS Exploit Kit of Mysterious Origin](https://cloud.google.com/blog/topics/threat-intelligence/coruna-powerful-ios-exploit-kit) (`daringfireball.net`)
  - TLDR: 名为 Coruna 的 iOS 漏洞利用工具包包含 23 个漏洞利用，影响 iOS 13-17.2.1，已被监控软件商、俄间谍组织及中国黑客多次使用，揭示活跃的零日漏洞二手市场。
  - WHAT: Coruna 是一个模块化 iOS 漏洞利用工具包，包含 5 条完整利用链和 23 个漏洞利用。其核心价值在于集成了针对多个 iOS 版本的、使用非公开技术及绕过缓解措施的先进漏洞利用，可被不同威胁行为者组合使用。
  - WHY: 1. 技术影响：展示了高级漏洞利用技术的扩散，单个工具包被不同国家背景和动机的行为者使用。2. 市场信号：证实存在活跃的零日漏洞“二手市场”，监控软件商可能是初始来源。3. 开发启示：凸显了即使是最新 iOS 版本也可能存在被集成到复杂工具包中的未知漏洞，对移动应用安全开发和第三方组件审计提出更高要求。
  - ACTION: 1. 立即行动：确保所有 iOS 设备已更新至最新版本（iOS 17.3+），以修复已知关联漏洞。2. 监控威胁：订阅 Google Threat Intelligence Group (GTIG) 等权威情报源警报。3. 代码审计：审查应用中集成的第三方 SDK 或库，特别是那些处理敏感数据或权限的组件。4. 纵深防御：在移动应用中实施运行时完整性检查、反调试和代码混淆，增加攻击成本。
  - TWEET: Coruna 工具包证明，顶级漏洞利用技术会像商品一样在暗网流通。你的应用是否依赖的某个 SDK 可能已内置了未公开的利用链？立即行动：1. 强制设备更新 2. 实施运行时保护 3. 订阅 GTIG 情报。安全不是功能，是基础设施。
