# OPMLWatch Digest

Generated at: 2026-02-18 17:21:54 UTC

- [The A.I. Disruption We’ve Been Waiting for Has Arrived](https://simonwillison.net/2026/Feb/18/the-ai-disruption/#atom-everything) (`simonwillison.net`)
  - TLDR: Claude Code 在 2025 年 11 月后能力质变，能独立完成复杂项目，颠覆传统软件开发成本模型。
  - WHAT: AI 编程工具（以 Claude Code 为代表）已能连续工作数小时，生成可用的完整网站与应用，将原本需数月、数十万美元的定制软件开发成本降至每月 200 美元的订阅费。
  - WHY: 大语言模型在代码生成、上下文理解与任务规划能力上取得突破，能处理数据迁移、网站重构等复杂、冗长的工程任务。
  - ACTION: 立即用 Claude Code 或同类工具复活个人积压的旧项目，并量化其替代传统外包工作的经济价值，重新评估自身技术债务与产品开发策略。
  - TWEET: 【颠覆性时刻】Paul Ford 亲历：Claude Code 在 11 月后“突然变强”，能独立开发完整应用。他仅用周末就完成了自己曾报价 2.5 万美元的网站重构，并处理了曾报价 35 万美元的数据集。AI 编程已从“助手”变为“执行者”，每月 200 美元即可创造数十万价值。你的旧项目该复活了。#AI编程 #开发者经济
- [Could Write­Process­Memory be made faster by avoiding the intermediate buffer?](https://devblogs.microsoft.com/oldnewthing/20260218-00/?p=112069) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: WriteProcessMemory 无法通过避免中间缓冲区来显著加速，因其设计初衷是调试器工具而非通用 IPC，强行优化会引入安全风险且收益极小。
  - WHAT: 博客探讨了通过内核 MDL 机制直接映射内存来消除 WriteProcessMemory 中间拷贝的可能性，但最终认为该方案过度复杂、不安全，且不符合函数实际使用场景。
  - WHY: WriteProcessMemory 的核心用途是调试器的小规模内存写入（如打补丁），并非大数据量 IPC；共享内存已是更优解。任何试图绕过拷贝的优化都可能将源进程的敏感页面暴露给目标进程的用户空间，造成严重安全漏洞。
  - ACTION: 进行进程间数据传输时，应优先选用共享内存机制。仅在调试等特定场景使用 WriteProcessMemory，并严格评估其安全限制，避免用于常规高性能通信。
  - TWEET: 【技术分析】WriteProcessMemory 能否跳过中间缓冲区加速？答案是否定的。其设计本为调试器小规模写入，非通用 IPC。强行用内核 MDL 优化会锁定页面、引入安全风险（可能泄露源进程内存），且收益微乎其微。高性能进程通信请直接用共享内存。详情：https://devblogs.microsoft.com/oldnewthing/20260218-00/?p=112069
- [Stream of Consciousness Driven Development](https://buttondown.com/hillelwayne/archive/stream-of-consciousness-driven-development/) (`buttondown.com/hillelwayne`)
  - TLDR: 通过创建结构化 Markdown 文件（问题-尝试-方案-理论）来记录并同步复杂决策的完整思考链，确保结对双方在修改前达成深度共识。
  - WHAT: 在修改规范或代码前，新建一个文档，按顺序写下：当前问题、问题细节、已尝试的失败方案、提出方案、方案背后的理论、具体实现方式、预期影响、未解决的问题等。
  - WHY: 解决因经验差距导致的理解偏差，避免一方因“不完全理解”而盲从专家，从而产生难以维护的产出。将隐性思考显性化，强制双方厘清逻辑。
  - ACTION: 下次在结对处理一个概念复杂的任务时，暂停编码/修改，花 15-30 分钟共同填写一个类似结构的 Markdown 文件，确认无误后再行动。
  - TWEET: 分享一个实用的结对编程技巧：面对复杂改动，先别急着动手。新建一个 Markdown 文件，把问题、试错、方案和理论都写下来，形成“思考链文档”。这能确保双方真正理解“为什么”，避免盲从，产出更健壮的代码/规范。试试看！#开发协作 #文档驱动 #工程实践
