# OPMLWatch Digest

Generated at: 2026-03-02 12:45:01 UTC

- [Adding "Log In With Mastodon" to Auth0](https://shkspr.mobi/blog/2026/03/adding-log-in-with-mastodon-to-auth0/) (`shkspr.mobi`)
  - TLDR: Auth0 官方不支持 Mastodon 登录，但可通过其 Connections API 结合 Mastodon 的动态应用注册功能，实现适配任意联邦宇宙实例的自定义 OIDC 身份提供商。
  - WHAT: 具体实现了一个自定义 Auth0 连接，该连接能动态为任意 Mastodon 实例注册应用（获取 client_id/secret），并作为 OIDC 身份提供商与 Auth0 集成，从而让用户能用其联邦宇宙账号登录。
  - WHY: 1. 满足 Fediverse 用户对去中心化社交登录的需求，提升产品包容性。2. 避免仅支持单一 Mastodon 实例的局限性，符合联邦宇宙的分布式特性。3. 绕过 Auth0 官方未提供的社交连接，自主控制身份验证流程。
  - ACTION: 1. 查阅 Auth0 Connections API 文档，了解自定义 OIDC 提供商配置。2. 参考开源代码（如 PHP 示例），实现动态向 Mastodon 实例注册应用并获取令牌的逻辑。3. 在 Auth0 仪表板创建新连接，填入动态生成的授权/令牌端点。4. 全面测试不同 Mastodon 实例的登录流程。
  - TWEET: 技术挑战：如何在中心化身份平台（Auth0）上支持去中心化社交网络（Fediverse）的登录？答案是自定义 OIDC 连接 + 动态客户端注册。详细解析与开源代码：
- [AMD Am386 released March 2, 1991](https://dfarq.homeip.net/amd-am386-released-march-2-1991/?utm_source=rss&utm_medium=rss&utm_campaign=amd-am386-released-march-2-1991) (`dfarq.homeip.net`)
  - TLDR: AMD Am386的发布比Intel 386晚了近6年，但这并非技术能力不足，而是其当时的战略选择与市场定位所致。
  - WHAT: Am386是AMD对Intel 80386的兼容克隆处理器，于1991年3月2日正式发布，最终在x86兼容机市场取得了成功。
  - WHY: 延迟主要因AMD当时将资源集中于其他业务（如闪存），并需应对与Intel的法律纠纷，同时等待合适的市场切入点。
  - ACTION: 研究处理器历史竞争案例，理解技术路线图决策如何受法律、市场与战略资源分配影响。
  - TWEET: AMD Am386的发布常被误读为技术落后。实则是战略选择：专注闪存、应对法律战、等待市场窗口。技术竞争不仅是工程，更是商业与法律的综合博弈。回顾历史，对理解当今芯片战仍有启示。
