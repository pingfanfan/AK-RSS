# OPMLWatch Digest

Generated at: 2026-03-20 22:40:58 UTC

- [Windows stack limit checking retrospective: arm64, also known as AArch64](https://devblogs.microsoft.com/oldnewthing/20260320-00/?p=112154) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 本文详解Windows在arm64架构下的栈限制检查机制，使用16字节对齐的'段落'作为单位，通过逐页探测确保栈空间安全，并与其他架构对比。
  - WHAT: chkstk是Windows用于验证栈空间的辅助函数。arm64版本以16字节对齐的'段落'为单位计算需求，通过向下取整到页面并探测来防止栈溢出，且不调整栈指针。
  - WHY: 理解栈限制检查对诊断栈溢出崩溃、优化大型栈分配性能以及深入系统安全机制至关重要，尤其在跨架构开发中。
  - ACTION: 检查代码中大型局部变量分配，了解chkstk在栈分配中的作用；参考文末表格，对比不同架构的栈检查差异。
  - TWEET: arm64栈检查用'段落'（16字节）作单位，巧妙支持大栈分配。页面探测机制防溢出，对比x86/MIPS等架构差异。底层机制，值得了解。
