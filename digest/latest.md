# OPMLWatch Digest

Generated at: 2026-03-16 14:49:06 UTC

- [How coding agents work](https://simonwillison.net/guides/agentic-engineering-patterns/how-coding-agents-work/#atom-everything) (`simonwillison.net`)
  - TLDR: 编码代理是LLM的封装，通过隐藏提示和可调用工具扩展其能力，核心在于理解LLM的token机制和提示工程。
  - WHAT: 编码代理是一种软件，作为大语言模型（LLM）的“ harness”，通过不可见的提示和可调用工具为其增加执行代码、查询文件等能力。
  - WHY: 理解其工作原理能帮助开发者更明智地选择和应用代理，避免因忽略token成本、上下文限制或多模态输入处理而导致的效率或安全问题。
  - ACTION: 1. 使用OpenAI Tokenizer实验文本如何转为token。2. 学习chat templated prompts结构。3. 试用Claude Code或Cursor等工具，观察其工具调用日志。
  - TWEET: 想用好AI编码代理？先搞懂LLM的token机制：1. 所有输入（文本/图片）都转成token计费 2. 上下文长度有限 3. 代理行为由系统提示词驱动。试试OpenAI tokenizer，你会对成本有新认识。
- [Windows stack limit checking retrospective: PowerPC](https://devblogs.microsoft.com/oldnewthing/20260316-00/?p=112140) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: PowerPC的Windows栈检查函数chkstk接受负的栈帧大小，是为了让调用者能使用原子指令stwxu高效创建栈帧。这体现了ISA特性对底层系统设计的直接影响。
  - WHAT: chkstk函数验证栈增长是否会导致越界。它接收一个负值参数（-帧大小），通过计算新栈指针并对比用户/内核栈限，决定是否需逐页探测（probe）以触发缺页。核心是处理PowerPC地址空间 midpoint 的用户/内核分割。
  - WHY: 1. 指令集限制：PowerPC的索引存储指令只有加法形式（add），无减法变体。使用负值参数配合stwxu（store with update indexed）能原子地完成‘新栈指针 = 旧栈指针 + 负值’的减法效果。2. 性能：原子操作避免竞态，确保栈创建在多线程环境下安全。3. 架构特性：利用PowerPC地址空间midpoint的硬性分割快速判断栈类型。
  - ACTION: 1. 阅读Windows Research Kernel或ReactOS源码中的chkstk实现，对比不同架构（x86, ARM, MIPS）的差异。2. 在编写平台相关汇编或JIT代码时，优先利用原子指令和架构特性优化关键路径。3. 关注ISA文档，理解寻址模式、原子性保证对系统设计的影响。
  - TWEET: 【技术深挖】Windows在PowerPC上的chkstk函数接受负的栈帧大小，核心是为了调用原子指令stwxu。这巧妙绕过了PowerPC索引存储只有加法的限制，实现了栈指针的原子更新。同时利用地址空间midpoint快速区分用户/内核栈。这种底层设计直接受ISA约束，是系统编程中‘硬件特性驱动软件设计’的经典案例。
