# OPMLWatch Digest

Generated at: 2026-03-10 17:44:26 UTC

- [LLMs are bad at vibing specifications](https://buttondown.com/hillelwayne/archive/llms-are-bad-at-vibing-specifications/) (`buttondown.com/hillelwayne`)
  - TLDR: AI能辅助形式化方法，但初学者直接生成 specs 常犯基础错误（如编译问题、错误抽象），需专家指导与严格验证。
  - WHAT: 本文通过分析一个由AI生成的Alloy规格案例，指出其在基础语法（缺失模块导入）和设计模式（误用布尔值而非子类型）上的错误，说明LLM缺乏对形式化规范深层语义的“直觉”。
  - WHY: 形式化方法的核心在于严谨的数学抽象，LLM基于统计模式生成代码，易产生表面正确但本质错误的spec，这可能导致安全关键系统的验证失效，阻碍形式化方法的可靠应用。
  - ACTION: 开发者应系统学习TLA+/Alloy基础，将AI视为辅助工具而非替代；生成spec后必须手动验证逻辑正确性、编译通过性，并参考权威模式；积极参与社区审查。
  - TWEET: 别让LLM的‘感觉’毁了你的形式化验证。分析一个AI生成的Alloy项目：基础语法错误+错误抽象设计。形式规格无容错空间，AI输出必须逐行审查。工具是助手，你才是专家。
