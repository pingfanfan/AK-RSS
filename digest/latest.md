# OPMLWatch Digest

Generated at: 2026-03-11 10:43:32 UTC

- [git-pkgs/actions](https://nesbitt.io/2026/03/11/git-pkgs-actions.html) (`nesbitt.io`)
  - TLDR: git-pkgs推出官方GitHub Actions，三行YAML即可在PR中自动展示依赖变更、扫描漏洞和检查许可证。
  - WHAT: 一套可复用的GitHub Actions，自动化git-pkgs工具链：setup初始化环境，diff展示依赖变更，vulns扫描OSV漏洞，licenses检查许可证。
  - WHY: 过去需手动下载二进制、初始化数据库并配置检查，现在通过标准化Actions，开发者能无缝集成到CI/CD，提升安全性和合规效率。
  - ACTION: 在项目.github/workflows/中添加三行YAML，使用git-pkgs/actions/setup和diff，立即在PR中查看依赖变更。
  - TWEET: git-pkgs/actions 上线：用 GitHub Actions 自动化依赖安全与合规检查。setup + diff + vulns + licenses，PR 集成超简单。https://nesbitt.io/2026/03/11/git-pkgs-actions.html
