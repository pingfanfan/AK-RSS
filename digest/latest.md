# OPMLWatch Digest

Generated at: 2026-03-11 19:43:01 UTC

- [Pluralistic: AI "journalists" prove that media bosses don't give a shit (11 Mar 2026)](https://pluralistic.net/2026/03/11/modal-dialog-a-palooza/) (`pluralistic.net`)
  - TLDR: 媒体公司用AI生成低质内容并强推自动播放视频，暴露其追求利润而牺牲用户体验与新闻伦理的本质，是技术滥用的典型案例。
  - WHAT: 文章批判媒体行业为降本增效，采用AI“记者”自动化生产内容，导致新闻质量下降、用户体验恶化（如强制自动播放），本质是资本对内容产业的侵蚀。
  - WHY: 对开发者警示：当技术决策仅服务于商业指标（点击率、成本）而忽视用户价值与伦理时，会制造垃圾产品。同时揭示内容平台基础设施如何被用于最大化剥削用户注意力。
  - ACTION: 开发AI内容工具时，内置伦理审查与用户体验控制（如默认禁用自动播放）；推动团队建立AI生成内容透明标注规范；拒绝为短期KPI优化而损害长期信任的技术方案。
  - TWEET: “模态对话框大杂烩”：当媒体老板把AI当廉价劳动力，连基本用户体验（如自动播放）都懒得管，就证明他们根本不在乎读者，只在乎流量变现。技术人，你的代码在助纣为虐吗？
- [How do compilers ensure that large stack allocations do not skip over the guard page?](https://devblogs.microsoft.com/oldnewthing/20260311-00/?p=112134) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 大栈分配时，编译器插入_chkstk调用，顺序触达所有跨页，确保保护页逐级下移，避免非法访问。
  - WHAT: 当函数栈帧超过一页且首个变量位于最低地址时，可能直接访问保留区而非保护页，导致进程崩溃。但实际不会，因为编译器有特殊机制。
  - WHY: 系统仅维护一个保护页（栈分配区下方）。若跳过，首次越界将访问未提交保留区而崩溃。_chkstk函数顺序访问所有跨页，触发系统将保护页转已提交、更新栈限并下移新保护页，确保首次越界总在保护页上。
  - ACTION: 开发者应了解此机制以调试栈溢出；编译器自动处理大栈分配，但知晓_chkstk原理有助于理解底层内存安全。建议阅读原文历史系列，深入x86栈管理演变。
  - TWEET: 大栈分配不跳过保护页的秘密：编译器插入_chkstk，顺序触达所有跨页，让保护页逐级下移。详解：https://devblogs.microsoft.com/oldnewthing/20260311-00/?p=112134
