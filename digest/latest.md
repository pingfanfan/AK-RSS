# OPMLWatch Digest

Generated at: 2026-02-15 07:42:18 UTC

- [How Michael Abrash doubled Quake framerate](https://fabiensanglard.net/quake_asm_optimizations/index.html) (`fabiensanglard.net`)
  - TLDR: Michael Abrash 通过极致汇编优化（内联汇编、循环展开、数据对齐）将 Quake 帧率提升 100%，核心是“用底层硬件控制换取性能”。
  - WHAT: 具体操作：1. 用内联汇编重写热点 C 代码；2. 循环展开减少跳转开销；3. 强制数据对齐提升缓存命中；4. 利用 Pentium 流水线特性调度指令。
  - WHY: 90 年代硬件限制严格，每周期指令数敏感。现代虽不同，但原理相通：减少分支预测失败、提升数据局部性、发挥 SIMD/并行能力，对游戏引擎、高频交易等仍关键。
  - ACTION: 1. 用 perf/VTune 定位热点；2. 在关键路径尝试内存布局优化或 SIMD 指令；3. 学习 Abrash《Graphics Programming Black Book》案例；4. 权衡可读性与性能，避免过度优化。
  - TWEET: 重温经典：Michael Abrash 如何用汇编将 Quake 帧率翻倍？核心手段：内联汇编、循环展开、数据对齐。启示：即使现代开发，热点代码仍需关注底层优化（缓存、分支、SIMD）。建议：1. Profiling 找瓶颈 2. 关键路径尝试内存/指令优化 3. 读《Graphics Programming Black Book》。#性能优化 #游戏开发 #底层编程
