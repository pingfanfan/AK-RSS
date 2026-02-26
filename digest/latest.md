# OPMLWatch Digest

Generated at: 2026-02-26 06:16:31 UTC

- [Google API Keys Weren't Secrets. But then Gemini Changed the Rules.](https://simonwillison.net/2026/Feb/26/google-api-keys/#atom-everything) (`simonwillison.net`)
- [‘H-Bomb: A Frank Lloyd Wright Typographic Mystery’](https://www.inconspicuous.info/p/h-bomb-a-frank-lloyd-wright-typographic) (`daringfireball.net`)
  - TLDR: 文章通过赖特建筑标识案例，揭示排版中常见的‘H/S’与‘P/Q’字形混淆问题，提醒开发者在UI/字体渲染中关注细节，避免因字形相似导致的可读性错误。
  - WHAT: 具体分析了在重新悬挂标识时，将衬线体字母‘H’和‘S’误认为无衬线体‘P’和‘Q’的经典排版错误。这涉及字体设计、字形历史及视觉感知，对前端字体选择、图标设计及安全关键界面有直接参考价值。
  - WHY: 对开发者而言，此案例凸显了UI细节的极端重要性：微小的字形混淆可能导致用户误解、操作错误，甚至安全风险。在AI生成内容、多语言渲染及无障碍设计中，字体选择与渲染引擎的细节必须经过严格测试。
  - ACTION: 1. 审计项目中的字体渲染，特别检查易混淆字符对（如H/S、P/Q、b/d）。2. 在设计系统中加入字形对比测试用例。3. 对团队进行排版基础知识培训，避免历史错误重演。
  - TWEET: 你遇到过字形混淆导致的UI问题吗？比如把‘H’看成‘P’？赖特这个案例提醒我们：字体渲染的细节绝不能马虎。分享你的经历👇 #开发者 #排版
- [Against Query Based Compilers](https://matklad.github.io/2026/02/25/against-query-based-compilers.html) (`matklad.github.io`)
- [Hyperbolic versions of latest posts](https://www.johndcook.com/blog/2026/02/25/hyperbolic-versions-of-latest-posts/) (`johndcook.com`)
  - TLDR: 通过双曲函数类比三角恒等式，并用Python代码验证，快速掌握数学转换规律。
  - WHAT: 本文展示了如何将三角函数恒等式（如|sin(x+iy)| = |sin x + sin iy|）及其反函数组合表，系统性地转换为对应的双曲函数形式，并提供Python代码片段验证转换正确性。
  - WHY: 帮助开发者深入理解数学结构，减少符号混淆；提供的验证代码能快速发现手算或实现中的错误，提升数值计算的可靠性。
  - ACTION: 运行提供的Python代码验证表中所有条目；尝试将其他已知三角恒等式（如和角公式）转换为双曲形式并验证。
  - TWEET: 三角恒等式有双曲‘镜像’？John D. Cook 展示如何系统转换，并用Python验证。代码虽不证明，但能抓错。开发者必看数学技巧。#编程 #数学
- [Members Only: Your anonymity set has collapsed and you don't know it yet](https://www.joanwestenberg.com/members-only-your-anonymity-set-has-collapsed-and-you-dont-know-it-yet/) (`joanwestenberg.com`)
- [The Last Gasps of the Rent Seeking Class](https://geohot.github.io//blog/jekyll/update/2026/02/26/the-last-gasps-of-the-rent-seeking-class.html) (`geohot.github.io`)
  - TLDR: AI正通过API民主化瓦解传统企业利用时间不对称的商业模式，芯片设计/软件层竞争开放无垄断风险，但应用层可能重现租金抽取。
  - WHAT: 分析AI如何终结基于人为摩擦（如呼叫中心、复杂流程）的租金经济，并评估AI供应链（电力、芯片制造、芯片设计/软件、模型、应用）五层的竞争格局与垄断风险。
  - WHY: 开发者需理解AI如何重构市场规则，识别基础设施层（如芯片设计、开放API）的真实机会与风险，避免在新生态中陷入封闭系统的租金陷阱。
  - ACTION: 1. 优先采用OpenRouter等开放API与模型；2. 评估芯片/模型层的开源项目（如RISC-V、Llama）；3. 在应用层构建时强调用户时间价值，避免添加非必要摩擦。
  - TWEET: AI终结摩擦经济？文章指出：企业靠耗尽用户时间盈利的时代将结束。关键在开放API与供应链分层——芯片设计/软件层无垄断，但应用层需警惕新租金抽取。开发者应拥抱开放生态。
- [Notes on Linear Algebra for Polynomials](https://eli.thegreenplace.net/2026/notes-on-linear-algebra-for-polynomials/) (`eli.thegreenplace.net`)
