# OPMLWatch Digest

Generated at: 2026-02-15 06:45:00 UTC

- [The problem of delivering errors to syndication feed readers](https://utcc.utoronto.ca/~cks/space/blog/web/FeedReaderErrorsProblem) (`utcc.utoronto.ca/~cks`)
  - TLDR: 针对被阻止的订阅源阅读器，返回HTML错误页或302重定向无效，必须返回有效的订阅源（如Atom/RSS）才能让用户看到解释。
  - WHAT: 订阅源阅读器期望响应为订阅源格式，若收到text/html等格式会静默丢弃，用户仅可能看到标题或状态码，精心编写的HTML错误页无法展示。
  - WHY: 订阅源阅读器协议设计如此，非订阅源内容被视为解析错误，这是其核心行为逻辑，而非缺陷。
  - ACTION: 1. 订阅源提供者：为被阻止的User-Agent等准备一个“错误订阅源”，返回200状态码，并在条目中嵌入解释或指向HTML页的链接。2. 订阅源阅读器开发者：尝试显示收到的HTML错误响应体，提升异常情况下的用户体验。
  - TWEET: 【订阅源错误处理指南】当需阻止某些阅读器时，返回HTML错误页无效，因其只解析RSS/Atom。正确做法：返回200状态码+一个含解释的特殊订阅条目。同时呼吁阅读器开发者：请尝试显示收到的HTML错误内容。这对提升订阅生态的健壮性至关重要。 #RSS #WebDev
