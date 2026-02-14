# OPMLWatch Digest

Generated at: 2026-02-14 10:41:07 UTC

- [Package Management Namespaces](https://nesbitt.io/2026/02/14/package-management-namespaces.html) (`nesbitt.io`)
  - TLDR: 扁平命名空间虽简洁，但导致好名字被占用、废弃包不回收，且易引发拼写错误攻击，威胁供应链安全。
  - WHAT: 博客以PyPI、crates.io等为例，分析扁平命名空间（全局先到先得）的优缺点：优点是命令短、易记忆；缺点是名字稀缺、废弃包滞留、易受typosquatting攻击。
  - WHY: 名字稀缺增加开发者认知负担与选择成本；拼写错误攻击是直接的安全威胁，影响基础设施的可靠性与开发者信任。
  - ACTION: 1. 项目优先选用带命名空间隔离的包管理器（如NPM、Go Modules）；2. 强制使用锁文件（如package-lock.json, Cargo.lock）并定期审计依赖；3. 团队内部分享命名规范，警惕相似包名。
  - TWEET: 包管理器命名空间设计至关重要。扁平命名空间（如PyPI）虽简洁，但导致好名字被占用、废弃包不回收，且易引发拼写错误攻击，严重威胁供应链安全。建议优先选择带命名空间隔离的系统（如NPM），项目中使用锁文件并定期审计依赖，团队内部分享命名规范，警惕相似包名。安全与可维护性需从基础设施设计开始。#开发者安全 #包管理
