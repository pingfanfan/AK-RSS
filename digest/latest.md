# OPMLWatch Digest

Generated at: 2026-03-03 18:44:33 UTC

- [Apple Debuts M5 Pro and M5 Max, and Renames Its M-Series CPU Cores](https://www.apple.com/newsroom/2026/03/apple-debuts-m5-pro-and-m5-max-to-supercharge-the-most-demanding-pro-workflows/) (`daringfireball.net`)
  - TLDR: 苹果推出M5 Pro/Max，采用双die融合架构，CPU核心重新分类为效率、超级、性能三类，多线程性能提升2.5倍。
  - WHAT: M5 Pro/Max采用全新Fusion Architecture，整合双die为单SoC。CPU为18核设计：6个超级核心（原性能核心，世界最快单线程）+12个新性能核心（能效优化）。所有M5芯片均含超级核心，基础M5仅有效率核心。
  - WHY: 命名调整旨在区分极致单线程性能（超级核心）与能效优先的多线程核心（新性能核心），以应对专业工作负载的多样化需求。架构整合提升带宽与能效。
  - ACTION: 开发者需更新性能分析工具以识别三类核心；评估现有应用在新核心调度下的能效比；关注苹果文档以适配命名变更。
  - TWEET: M5系列CPU核心大调整：效率、超级（原性能）、性能（新）三类。M5 Pro/Max的18核（6超+12新性能）多线程性能飙升2.5倍。双die融合架构带来更高带宽，专业用户注意核心调度优化。 #开发者 #硬件
- [Pluralistic: Supreme Court saves artists from AI (03 Mar 2026)](https://pluralistic.net/2026/03/03/its-a-trap-2/) (`pluralistic.net`)
  - TLDR: 最高法院驳回AI公司请愿，确立‘创作时固定’版权原则，艺术家训练数据版权获保护。
  - WHAT: 美国最高法院驳回了AI公司关于使用受版权保护作品训练大模型的请愿，维持了下级法院判决，确认版权‘在创作固定时产生’，AI公司未经许可使用作品构成侵权。
  - WHY: 此案为AI训练数据使用设立了关键法律先例。开发者需重新评估数据管道合规性，确保训练数据有合法授权，否则将面临侵权风险。长期看，可能推动合成数据或授权数据集的发展。
  - ACTION: 立即审计训练数据版权；建立数据来源审查机制；关注版权法修订动态；评估合成数据或授权数据集方案。
  - TWEET: 最高法院站队艺术家！AI训练版权案尘埃落定，‘创作时固定’原则成关键。你的模型训练数据干净吗？#技术趋势 #AI安全
