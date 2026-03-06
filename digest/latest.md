# OPMLWatch Digest

Generated at: 2026-03-06 06:09:44 UTC

- [Agentic manual testing](https://simonwillison.net/guides/agentic-engineering-patterns/agentic-manual-testing/#atom-everything) (`simonwillison.net`)
  - TLDR: 编码代理虽能自执行代码，但自动化测试仍有盲区，必须结合手动测试（如python -c或/tmp临时文件）确保功能完整。
  - WHAT: 一种工程模式：指导AI代理在编写代码后，通过直接执行（如命令行运行、临时文件）进行手动验证，补充自动化测试的不足。
  - WHY: 自动化测试可能遗漏UI显示、启动崩溃等实际问题；手动测试提供真实环境反馈，对API和边缘案例至关重要，是质量保障的最后防线。
  - ACTION: 在提示词中明确要求代理使用`python -c`或`/tmp`临时文件进行手动测试，并人工检查执行结果，尤其关注边缘案例和API响应。
  - TWEET: AI编码代理能写代码，但别信它自动通过测试就完事。手动测试才是王道：用`python -c`或`/tmp`临时跑一跑，揪出自动化测试漏掉的坑。安全、质量双保障。#AI开发 #工程实践
