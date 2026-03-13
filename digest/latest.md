# OPMLWatch Digest

Generated at: 2026-03-13 06:11:39 UTC

- [Shopify/liquid: Performance: 53% faster parse+render, 61% fewer allocations](https://simonwillison.net/2026/Mar/13/liquid/#atom-everything) (`simonwillison.net`)
  - TLDR: Shopify CEO 使用 AI 代理对 Liquid 模板引擎进行 120 次实验，实现解析渲染性能提升 53%，内存分配减少 61%。关键优化包括字节级扫描和整数缓存。
  - WHAT: 这是 Shopify 对其开源 Ruby 模板引擎 Liquid 的性能优化 PR，由 CEO Tobias Lütke 使用 Karpathy 的 autoresearch AI 代理完成，通过 93 次提交、120 次实验达成。
  - WHY: 展示了 AI 代理在成熟代码库中进行微优化的潜力，凸显了健壮测试套件对自动化实验的关键作用，为其他项目提供了可复现的性能优化路径。
  - ACTION: 为你的项目建立完善的测试套件；尝试用 AI 代理进行性能微调；关注字节级操作和缓存策略。
  - TWEET: Shopify Liquid 性能优化 PR 解析：53% 速度提升来自字节扫描(String#byteindex)替代 StringScanner，纯字节 parse_tag_token 避免重置，以及 0-999 整数 to_s 缓存。120 次自动化实验证明：健壮测试套件 + AI 代理 = 高效微调。 #Ruby #Performance
- [Accents](https://mahdi.jp/apps/accents) (`daringfireball.net`)
- [The web in 1000 lines of C](https://maurycyz.com/projects/tinyweb/) (`maurycyz.com`)
  - TLDR: 用约1000行C代码构建极简浏览器，仅支持HTTP/HTML基础渲染，揭示Web核心本质。
  - WHAT: 一个约1000行C代码的微型浏览器项目，可发起HTTP请求、解析HTML并渲染纯文本页面，支持链接点击。
  - WHY: 批判现代浏览器（如Chromium）的4900万行代码复杂性，证明Web基础功能可极简实现，适合开发者理解HTTP/HTML核心。
  - ACTION: 访问项目链接，阅读代码；尝试用自己语言实现极简浏览器；分享你对Web简化的看法。
  - TWEET: 千行C代码重构Web浏览：一个极简浏览器项目，挑战Chromium的4900万行代码。值得每个开发者深思。链接：https://maurycyz.com/projects/tinyweb/
- [Pluralistic: Three more AI psychoses (12 Mar 2026)](https://pluralistic.net/2026/03/12/normal-technology/) (`pluralistic.net`)
  - TLDR: 文章批判当前AI领域盛行的三种非理性狂热：技术决定论、亿万富翁唯我论和纯粹主义文化，警告这些‘精神病’正在制造技术泡沫并扭曲投资与开发方向。
  - WHAT: 1. 技术决定论：认为AI发展是不可避免的线性进程，忽视社会与伦理影响。2. 亿万富翁唯我论：富豪将个人幻想投射到AI，用资本强行塑造未来。3. 纯粹主义文化：追求‘纯净’技术栈，排斥实用妥协，导致工程脱离实际需求。
  - WHY: 这些‘精神病’导致资本错配、资源浪费于华而不实的概念，挤压真正解决安全、基础设施等基础问题的空间。它们制造了脱离现实的期望，最终可能引发类似历史泡沫的崩溃，伤害开发者与用户。
  - ACTION: 作为开发者，应保持批判性思维，优先关注可验证的客户价值与系统可靠性。在技术选型中拒绝纯粹主义，拥抱务实工程。警惕被资本炒作的‘颠覆性’叙事绑架，专注于构建安全、可持续的产品。
  - TWEET: AI炒作新高峰？三种‘精神病’正在主导行业：1. 技术宿命论 2. 富豪幻想工程 3. 技术栈纯粹主义。结果？资本错配，基础安全被忽视。开发者该如何应对？保持清醒，用代码解决真问题，而非追逐泡沫。
- [Everything's Casino](https://www.joanwestenberg.com/everythings-casino/) (`joanwestenberg.com`)
