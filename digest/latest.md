# OPMLWatch Digest

Generated at: 2026-03-08 10:39:29 UTC

- [If It Quacks Like a Package Manager](https://nesbitt.io/2026/03/08/if-it-quacks-like-a-package-manager.html) (`nesbitt.io`)
  - TLDR: 本文揭示GitHub Actions等工具已演变为隐性包管理器，存在传递依赖风险却缺乏锁文件等防护，开发者需警惕供应链安全。
  - WHAT: 文章分析了GitHub Actions如何通过动作复用形成传递依赖树，尽管其设计初衷并非包管理器，但已具备相关特征并暴露安全短板。
  - WHY: 传递依赖引入未审计代码执行风险，而无锁文件机制使复现和审计变得困难，威胁CI/CD管道安全。
  - ACTION: 立即对GitHub Actions使用SHA固定，定期审计动作依赖树，并推动平台提供锁文件支持以增强可追溯性。
  - TWEET: GitHub Actions实为‘不完整包管理器’：递归下载动作却无锁文件，传递依赖风险高。立即SHA固定并审计依赖树，守护CI/CD安全。
