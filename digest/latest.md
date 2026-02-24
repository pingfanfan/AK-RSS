# OPMLWatch Digest

Generated at: 2026-02-24 16:18:18 UTC

- [go-size-analyzer](https://simonwillison.net/2026/Feb/24/go-size-analyzer/#atom-everything) (`simonwillison.net`)
  - TLDR: go-size-analyzer 是一个可视化工具，用于分析 Go 二进制文件的依赖大小分布，支持本地运行和浏览器在线分析。
  - WHAT: 这是一个基于树状图（treemap）的 Go 二进制文件大小分析工具。它能将编译后的二进制文件分解为标准库、主包、生成代码及未知节（如调试信息）等组成部分，并直观展示各部分的体积占比。
  - WHY: 理解二进制体积构成是优化应用大小、减少攻击面、提升部署效率的关键。该工具能快速定位体积“大户”（如未使用的标准库包、庞大的调试信息），帮助开发者在安全与基础设施层面做出数据驱动的决策。
  - ACTION: 立即访问在线版本（gsa.zxilly.dev）上传你的 Go 二进制文件进行分析。重点关注“Std Packages”和“Unknown Sections”（尤其是 DWARF 调试信息），评估是否可通过编译标志（如 -ldflags="-s -w"）或依赖裁剪来优化。
  - TWEET: Go 二进制大小分析神器 go-size-analyzer，树状图直观展示依赖体积，在线版即开即用。优化二进制体积，从理解构成开始。https://gsa.zxilly.dev
