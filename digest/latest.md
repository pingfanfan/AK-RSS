# OPMLWatch Digest

Generated at: 2026-03-04 17:43:20 UTC

- [Anti-patterns: things to avoid](https://simonwillison.net/guides/agentic-engineering-patterns/anti-patterns/#atom-everything) (`simonwillison.net`)
  - TLDR: 在AI辅助编程中，提交PR前必须亲自验证代码功能，避免将未审查的代码负担转嫁给同事。
  - WHAT: 指开发者使用AI生成代码后，不进行自身审查和验证，就直接提交大型PR，将审查和调试工作完全推给协作者的反模式。
  - WHY: 这浪费协作者时间、降低团队效率，且违背开发者对代码质量负责的核心职责。你未提供任何超越AI本身的额外价值。
  - ACTION: 1. 亲自运行并测试代码，确保其工作。2. 拆分大PR为多个小、聚焦的变更。3. 在PR描述中提供手动测试笔记、截图或视频作为证据。4. 审查并验证AI生成的PR描述文本。
  - TWEET: 用AI写代码后，自己先跑通再提PR。把未验证的代码扔给同事审查，是最低效且不负责任的协作方式。 #开发者 #AI辅助
