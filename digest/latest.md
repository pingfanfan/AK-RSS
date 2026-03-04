# OPMLWatch Digest

Generated at: 2026-03-04 18:44:54 UTC

- [Aha, I found a counterexample to the documentation that says that Query­Performance­Counter never fails](https://devblogs.microsoft.com/oldnewthing/20260304-00/?p=112110) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 文档称 QueryPerformanceCounter 在 WinXP+ 永不失败，但若传入未 8 字节对齐的 LARGE_INTEGER 指针，它会因 ERROR_NOACCESS 失败。这是参数无效，非 API 缺陷。
  - WHAT: 有开发者发现，在 Windows XP 及更高版本上，QueryPerformanceCounter 确实可能失败。方法是传入一个未正确对齐（非 8 字节）的 LARGE_INTEGER 结构体指针。文档的“永不失败”声明，建立在传入参数（包括指针指向的内存）完全有效且符合对齐要求的前提下。
  - WHY: LARGE_INTEGER 结构体在 Windows 上要求 8 字节对齐（默认 /Zp8 打包）。如果结构体实例本身在栈上未对齐，或通过 malloc 等分配的内存未保证对齐，传入的指针就无效。这违反了 API 的基本前置条件，导致函数行为未定义——可能失败、返回脏数据或崩溃。
  - ACTION: 1. 声明 LARGE_INTEGER 变量时，依赖编译器默认对齐通常安全。2. 动态分配时，使用 _aligned_malloc 或 C++17 aligned_alloc 确保 8 字节对齐。3. 审查旧代码或第三方库中自定义的时间戳结构体，检查其对齐属性。4. 理解文档隐含前提：你必须提供有效的参数。
  - TWEET: 有趣的反例：QueryPerformanceCounter 文档说它永不失败，但有人用未对齐的 LARGE_INTEGER 指针让它返回了 ERROR_NOACCESS。这提醒我们，API 文档的“永不失败”常隐含“参数必须有效”的前提。内存对齐这种底层细节，依然是 Windows 系统编程的必修课。
