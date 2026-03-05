# OPMLWatch Digest

Generated at: 2026-03-05 10:43:36 UTC

- [Package Manager Magic Files](https://nesbitt.io/2026/03/05/package-manager-magic-files.html) (`nesbitt.io`)
  - TLDR: 除manifest和lockfile外，包管理器还有大量隐藏配置文件控制注册表、认证和安装行为。其中.npmrc的script-shell设置可被利用劫持脚本执行，构成供应链安全风险。
  - WHAT: 本文深入探讨npm、Yarn等包管理器中常被忽视的配置文件，如.npmrc、.yarnrc.yml。它们管理注册表URL、认证令牌、缓存行为，甚至可指定任意可执行文件在生命周期脚本中运行，直接影响依赖安装的安全性和一致性。
  - WHY: 这些文件文档不全、命名不一，开发者易在不知情下提交敏感配置（如私有注册表令牌）或引入恶意设置。攻击者可利用它们进行供应链攻击，如通过.npmrc的script-shell执行任意代码，而无需修改package.json。
  - ACTION: 1. 使用git历史扫描工具检查是否误提交了.npmrc等敏感文件；2. 在CI/CD流水线中添加步骤，验证配置未被篡改；3. 团队内部分享包管理器配置安全最佳实践；4. 阅读所用包管理器的完整配置文档，了解所有可用选项。
  - TWEET: 包管理器的隐藏配置：.npmrc的script-shell设置可指定任意shell执行脚本，风险极高。确保不将其提交至仓库，并定期审计配置。安全无小事。 #DevOps #Infosec
