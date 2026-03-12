# OPMLWatch Digest

Generated at: 2026-03-12 22:39:59 UTC

- [Windows stack limit checking retrospective: x86-32, also known as i386](https://devblogs.microsoft.com/oldnewthing/20260312-00/?p=112136) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 原始Windows _chkstk函数通过逐页触摸内存来安全分配大栈帧，并采用一种‘被调用者弄脏栈’的非常规约定返回，这源于16位时代的兼容性设计。
  - WHAT: _chkstk函数接收EAX中的分配大小，模拟栈指针递减并逐页写入零以提交物理页面，最终将真实ESP调整到模拟位置，并直接跳转到调用者的返回地址。这是一种既非调用者清理也非被调用者清理的‘被调用者弄脏’约定。
  - WHY: 这种设计最初来自16位8086的chkstk，用于在远/近调用不同场景下更新栈限制并检查溢出。在32位下保留是为了兼容性和最小化生成代码的改动，尽管它引入了极其特殊的调用约定，增加了理解和调试复杂度。
  - ACTION: 在分析旧系统或调试底层问题时，注意这种非常规约定。现代编译器已不再使用此原始实现，但理解其原理有助于掌握栈提交机制和 historic 性能考量。
  - TWEET: 深入旧文：Windows x86-32的栈检查函数_chkstk如何工作？它不按常理出牌——被调用者修改栈指针并直接跳回，形成‘callee-dirty’约定。这源于8086时代，是历史与兼容性的有趣案例。
