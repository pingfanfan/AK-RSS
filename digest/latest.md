# OPMLWatch Digest

Generated at: 2026-02-26 20:43:26 UTC

- [How to Securely Erase an old Hard Drive on macOS Tahoe](https://www.jeffgeerling.com/blog/2026/securely-erase-hard-drive-macos-tahoe/) (`jeffgeerling.com`)
  - TLDR: 在 macOS Tahoe 的磁盘工具中，Apple 移除了机械硬盘的安全擦除选项。开发者需改用命令行工具如 diskutil 实现安全擦除。
  - WHAT: 问题：通过 USB 连接旧机械硬盘后，磁盘工具中缺少 'Security Options' 按钮，无法进行安全擦除。
  - WHY: 原因：Apple 可能认为现代 Mac 已不再使用机械硬盘，因此简化了磁盘工具，移除了相关功能。
  - ACTION: 解决方案：使用终端命令。全盘安全擦除：diskutil secureErase 0 /dev/diskX；仅擦除空闲空间：diskutil secureErase freespace 0 /Volumes/diskX。请先通过 diskutil list 确认磁盘标识符。
  - TWEET: 【macOS Tahoe 安全擦除陷阱】磁盘工具不再提供机械硬盘安全擦除选项！改用命令行：diskutil secureErase。开发者注意数据残留风险。 #macOS #安全擦除 #开发者
- [Hoard things you know how to do](https://simonwillison.net/guides/agentic-engineering-patterns/hoard-things-you-know-how-to-do/#atom-everything) (`simonwillison.net`)
  - TLDR: 主动积累已验证的技术解决方案（代码、工具、报告），这不仅能提升你解决复杂问题的能力，更能为AI代理提供高质量的‘积木’，驱动创新。
  - WHAT: 指系统性地收集和整理你亲自验证过的、能解决具体问题的技术实现。例如：一个单文件HTML工具、一个GitHub上的概念验证仓库、一份AI辅助生成的研究报告。核心是‘可运行的代码’和‘第一手经验’。
  - WHY: 1. 深化个人技术视野：知道‘什么可能’与‘亲眼见过实现’有巨大差异，积累能培养精准的技术判断力。2. 赋能AI代理：这些积累是极佳的提示词上下文，让AI能基于‘已知可行’的方案进行组合创新，而非从零幻想。
  - ACTION: 1. 立即开始记录：为每个解决过的问题创建一个最小化、可运行的示例（推荐单HTML文件或小型GitHub仓库）。2. 建立索引：用博客、TIL或笔记工具为这些‘技术积木’打上标签（如#OCR #内存优化）。3. 主动组合：向AI代理提出请求，要求它基于你库中2-3个现有示例，组合生成新方案。
  - TWEET: Simon Willison 的核心建议：把‘我知道怎么做到X’变成可运行的代码存档。这既是个人知识资产，也是喂给AI代理的优质数据。你的技术‘囤积’越丰富，AI帮你组合创新的空间就越大。
