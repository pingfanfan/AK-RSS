# OPMLWatch Digest

Generated at: 2026-03-17 22:06:54 UTC

- [Quoting Ken Jin](https://simonwillison.net/2026/Mar/17/ken-jin/#atom-everything) (`simonwillison.net`)
  - TLDR: CPython 3.15 alpha的JIT编译器在macOS AArch64上比尾调用解释器快11-12%，在x86_64 Linux上快5-6%，性能目标提前达成。
  - WHAT: CPython JIT是Python 3.15引入的即时编译器，通过将热点代码编译为机器码来提升执行速度，本次测试验证了其在主流平台的有效性。
  - WHY: 性能提升直接影响AI训练、数据处理等计算密集型任务的效率，降低基础设施成本，是Python核心优化的重要里程碑。
  - ACTION: 立即下载Python 3.15 alpha版本，对关键代码路径进行基准测试，评估JIT在您的 workload 中的实际收益。
  - TWEET: 🚀 Python 3.15 JIT性能提前达标！macOS ARM快11-12%，Linux x86快5-6%。开发者们，是时候测试你的代码了！#Python #Performance
