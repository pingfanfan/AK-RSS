# OPMLWatch Digest

Generated at: 2026-03-17 14:01:00 UTC

- [Samsung Discontinues Its Galaxy Z TriFold After Just Three Months](https://www.theverge.com/tech/895879/samsung-galaxy-z-trifold-discontinued-stock-sold-out) (`daringfireball.net`)
- [Windows stack limit checking retrospective: x86-32 also known as i386, second try](https://devblogs.microsoft.com/oldnewthing/20260317-00/?p=112144) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 新_chkstk通过ret指令避免返回地址预测器不同步，但因此与影子栈（CET）不兼容，且保持特殊的调用约定。
  - WHAT: Windows在x86-32上用于大栈分配时探测页面的_chkstk函数的新实现，通过复制返回地址并执行ret来调整栈指针。
  - WHY: 为避免返回地址预测器不同步导致的性能损失，采用ret指令直接返回，但破坏了影子栈的控制流完整性，且ABI已固化此行为。
  - ACTION: 开发者需了解此ABI变更，若使用CET影子栈需避免混合新旧chkstk，或考虑其他防护方案；测试链接兼容性。
  - TWEET: Windows栈检查新方案：ret指令提升性能，但牺牲CET影子栈兼容。ABI变更已固化，开发者需警惕混合编译风险。详情：https://devblogs.microsoft.com/oldnewthing/20260317-00/?p=112144
- [Powers don’t clear fractions](https://www.johndcook.com/blog/2026/03/17/powers-dont-clear-fractions/) (`johndcook.com`)
- [Help I'm being persecuted](https://www.experimental-history.com/p/help-im-being-persecuted) (`experimental-history.com`)
