# OPMLWatch Digest

Generated at: 2026-03-16 06:25:32 UTC

- [What is agentic engineering?](https://simonwillison.net/guides/agentic-engineering-patterns/what-is-agentic-engineering/#atom-everything) (`simonwillison.net`)
- [CHM Live: Apple at 50](https://www.youtube.com/live/eCSNJgI2LFI) (`daringfireball.net`)
- [Twelve-tone composition](https://www.johndcook.com/blog/2026/03/15/twelve-tone-composition/) (`johndcook.com`)
  - TLDR: 十二音技法通过强制使用12音序列的排列来避免调性，其逆行、倒置等操作构成Z₂×Z₂群，展示了数学规则如何约束音乐创作。
  - WHAT: 一种无调性作曲方法：作曲家创建一个12个半音的序列（音列），并严格按此顺序使用音符（可移调、变节奏）。通过逆行（R）、倒置（I）等操作生成变体，形成4种基本形式（P, R, I, RI）。
  - WHY: 对开发者而言，这是算法约束创造力的典型案例：用有限规则（12音序列+群操作）避免模式化，类似密码学置换或程序化生成。作者认为数学应用于节奏比旋律更易听，提示规则需适配人类感知。
  - ACTION: 尝试用代码生成12音序列（如随机排列），并实现R/I/RI操作；探索其他音乐数学结构（如节奏的欧几里得算法）；思考如何将此类约束系统应用于UI设计或安全协议。
  - TWEET: 十二音序列的数学：12!种可能，循环等价下剩11!种。逆行/倒置生成阿贝尔群Z₂×Z₂。规则越严，越需创造力——这对算法设计有何启示？
- [Shower Thought: Git Teleportation](https://idiallo.com/byte-size/git-teleportation?src=feed) (`idiallo.com`)
- [Food, Software, and Trade-offs](https://blog.jim-nielsen.com/2026/food-software-and-trade-offs/) (`blog.jim-nielsen.com`)
  - TLDR: 用食物类比阐述软件选择本质是权衡：AI生成代码与人工代码有根本区别，不存在普适“最佳”，需根据场景评估得失。
  - WHAT: 通过麦当劳派、市售派和自制派的例子，讨论软件中“最佳”定义的多样性，强调所有方案都有权衡，AI生成软件与人工软件不可简单等同。
  - WHY: 帮助开发者在AI时代避免非黑即白思维，理性评估技术选型，理解不同方案的适用场景与代价，做出更明智的架构与产品决策。
  - ACTION: 在技术决策中明确核心指标（如性能、安全、成本），为每个选项列出“得”与“失”，避免追求“全能”方案，定期复盘权衡是否仍符合目标。
  - TWEET: 软件选型如同选食物：AI生成快但可能牺牲可审计性，人工编写慢但更可控。关键不是谁更好，而是你的场景需要放弃什么。 #技术权衡
