# OPMLWatch Digest

Generated at: 2026-02-23 07:21:38 UTC

- [Red/green TDD](https://simonwillison.net/guides/agentic-engineering-patterns/red-green-tdd/#atom-everything) (`simonwillison.net`)
  - TLDR: 通过‘红绿TDD’模式指导AI编码代理：先写自动化测试并确认其失败（红），再编写代码使测试通过（绿）。这能有效防止AI生成无效或冗余代码，并自动构建回归测试套件。
  - WHAT: 红绿TDD是测试驱动开发（TDD）的严格实践。开发者必须先编写自动化测试，观察测试失败（红阶段），然后迭代实现代码直至测试通过（绿阶段）。这确保了每段代码都有对应测试，且测试真正验证了功能。
  - WHY: 对AI编码代理而言，红绿TDD至关重要。它直接对抗AI生成无效代码或构建未使用功能的常见风险，同时强制建立全面的自动化测试套件，为项目增长提供回归防护，是保持代码库健康的最有效手段。
  - ACTION: 在向AI编码代理（如Claude Code、Cursor）发出提示时，明确包含‘使用红绿TDD’或‘先写测试’。要求代理执行测试并展示红绿阶段结果，验证测试失败后再实现代码。
  - TWEET: 红绿TDD是AI编码代理的‘黄金法则’。先写测试并确认失败，再写代码通过测试。这能杜绝AI的幻觉代码，并生成可维护的测试套件。强烈建议在提示词中强制使用此模式。#编程 #人工智能
