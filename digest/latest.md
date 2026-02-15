# OPMLWatch Digest

Generated at: 2026-02-15 06:13:50 UTC

- [How Generative and Agentic AI Shift Concern from Technical Debt to Cognitive Debt](https://simonwillison.net/2026/Feb/15/cognitive-debt/#atom-everything) (`simonwillison.net`)
  - TLDR: AI代理加速开发，但团队对系统的共享理解会快速流失，导致“认知债务”累积，最终使项目陷入瘫痪。
  - WHAT: 认知债务指开发者因快速迭代（尤其是AI生成代码）而丧失对系统设计意图和整体架构的心智模型，即使代码可读，团队也无法有效修改。
  - WHY: 生成式AI让代码产出速度远超人类理解速度，债务从“代码库”转移到“人脑”，更隐蔽且直接削弱团队持续交付能力。
  - ACTION: 1. 强制AI生成代码后撰写设计注释；2. 每周举行15分钟设计意图同步会；3. 用架构决策记录（ADR）固化关键决策；4. 代码审查必须包含“是否易理解”检查项。
  - TWEET: 警惕！用AI写代码可能让你更快失去对项目的掌控。最新研究指出，最大的风险不是技术债务，而是“认知债务”——团队集体遗忘系统为何如此设计。即使代码整洁，无人能解释时，项目即瘫痪。对策：强制AI生成设计注释、定期同步意图、用ADR记录决策。别让速度吞噬理解。#认知债务 #AI编程 #工程实践
- [Two different tricks for fast LLM inference](https://seangoedecke.com/fast-llm-inference/) (`seangoedecke.com`)
  - TLDR: Anthropic与OpenAI的“fast mode”本质不同：前者通过低批次推理加速真实模型Opus 4.6（2.5倍速，成本高），后者依赖专用芯片运行简化模型GPT-5.3-Codex-Spark（15倍速，能力降级）。
  - WHAT: Anthropic快模式：真实Opus 4.6，吞吐约170 tokens/秒（原65），价格约6倍。OpenAI快模式：简化模型Spark，超1000 tokens/秒（原65），但工具调用等能力明显弱于完整版GPT-5.3-Codex。
  - WHY: 核心差异在技术路径：Anthropic采用“低批次”策略，减少用户等待批处理填充时间以提升单请求速度；OpenAI则可能使用Cerebras等专用芯片实现硬件级加速，但以模型能力简化为代价。
  - ACTION: 开发者需按场景选择：若任务依赖复杂工具调用或高精度，选Anthropic真实模型并接受成本/速度权衡；若为高吞吐、低复杂度任务，可试OpenAI快模式。自建服务时，可通过调低批次大小模拟Anthropic模式以降低延迟。
  - TWEET: Anthropic与OpenAI的“fast mode”大不同：一个用真实模型Opus 4.6（低批次推理，快2.5倍但贵6倍），一个用简化模型Spark（专用芯片，快15倍但能力降级）。选型关键：任务是否需要完整模型能力？自建服务可调低批次优化延迟。 #AI #开发者
- [Wagon’s algorithm in Python](https://www.johndcook.com/blog/2026/02/14/wagons-algorithm-in-python/) (`johndcook.com`)
  - TLDR: 本文完整实现了Wagon算法，用数论方法高效求解奇素数p = a² + b²，核心是找二次非剩余、计算-1的平方根，并用修改版欧几里得算法，关键是用`isqrt`避免大数浮点精度问题。
  - WHAT: 博客提供了可直接运行的Python代码，包含`find_nonresidue`（找二次非剩余）、`my_euclidean_algorithm`（修改欧几里得算法，以`isqrt(p)`为停止条件）和`find_ab`（主函数），并演示了对大素数p=2²⁵⁵-19的求解。
  - WHY: 对开发者而言，这是将抽象数论（二次剩余、欧几里得算法）转化为高效、可处理极大整数（如密码学中常见的大素数）的实用代码范例，解决了传统公式法在大数下不实用的问题。
  - ACTION: 直接复制文末代码到你的Python环境（需sympy），用`find_ab(p)`求解任意满足p≡1 mod 4的奇素数。重点理解`isqrt`替代`sqrt`处理大整数的必要性，并掌握算法三步：找非剩余、求平方根、执行修改欧几里得算法。
  - TWEET: 如何用Python将大奇素数p拆成两个平方数之和？John Cook博客详解Wagon算法：1. 找二次非剩余；2. 计算-1模p的平方根；3. 对(p, x)运行修改版欧几里得算法，停止条件为`isqrt(p)`。代码已给出，完美解决大数浮点精度问题，适用于密码学等场景。 #Python #数论 #密码学 #算法
- [The empire always falls](https://www.joanwestenberg.com/the-empire-always-falls/) (`joanwestenberg.com`)
  - TLDR: 警惕将当前AI巨头的统治视为永恒，历史表明所有技术-商业帝国都会因内部腐化而崩溃。
  - WHAT: 博客以罗马帝国、托勒密模型等历史案例，论证“永久性”是危险幻觉，并批评AI领域对OpenAI等公司的“必然性”叙事，指出系统内部脆弱性常被忽视。
  - WHY: 对开发者而言，盲目追随单一技术栈或巨头平台存在巨大风险，可能面临技术锁死、伦理失控或范式突变导致的系统性崩溃。
  - ACTION: 1. 立即评估核心业务对特定AI供应商的依赖，制定数据与接口迁移预案；2. 主动关注并测试开源/去中心化AI模型（如Llama、Mistral）作为备选；3. 在技术选型中优先考虑可解释性、可控性与合规性，而非单纯追求性能指标。
  - TWEET: 警惕AI“必然性”迷思！博客以罗马帝国、托勒密体系为例，指出所有技术帝国都终将崩溃。对开发者：别死磕单一AI巨头，评估风险，拥抱开源替代方案，为范式转移做准备。历史教训：永久性只是幻觉。#AI #开发者 #风险
- [tiny corp’s product – a training box](https://geohot.github.io//blog/jekyll/update/2026/02/15/tiny-corp-product.html) (`geohot.github.io`)
  - TLDR: 博客核心观点：当前LLM同质化趋势下，若仅依赖prompt个性化，云服务将胜出；本地模型唯一出路是支持“全量学习”以实现深度定制。tiny corp正通过tinybox硬件和tinygrad软件为此布局。
  - WHAT: 作者Geohot分析AI未来分歧：模型权重冻结、学习限于上下文时，云因成本优势主导；本地部署需提供每用户/组织的完全学习能力。公司已销售tinybox，并持续优化tinygrad以降低本地训练门槛。
  - WHY: 因闭源模型（如Claude）的个性化仅靠10kB的CLAUDE.md（prompt）实现，本质是“costuming”而非学习，将导致多样性崩溃与用户边缘化。本地模型必须突破此限制，让独特性真正融入权重。
  - ACTION: 1. 评估tinybox等硬件是否满足本地微调需求；2. 在项目中实践LoRA等高效微调技术，而非仅依赖prompt；3. 关注tinygrad进展，测试其训练性能；4. 设计产品时考虑用户数据所有权与模型定制权。
  - TWEET: Geohot新博文预警：若AI只靠prompt“伪装”个性，云将垄断市场！本地模型必须支持深度微调才有未来。tiny corp已推出tinybox硬件备战。开发者应警惕同质化，探索可定制模型方案，避免沦为“克隆人”。全文：https://geohot.github.io//blog/jekyll/update/2026/02/15/tiny-corp-product.html #AI #开发者 #基础设施
