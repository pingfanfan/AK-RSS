# OPMLWatch Digest

Generated at: 2026-03-02 10:43:58 UTC

- [Transitive Trust](https://nesbitt.io/2026/03/02/transitive-trust.html) (`nesbitt.io`)
  - TLDR: 开源依赖的信任链脆弱：你不仅信任直接依赖，还信任其维护者及其所有依赖的维护者，形成无限传递。工具只能扫描表层，深层构建依赖常被忽视。
  - WHAT: 指在依赖管理中，你间接信任所有上游维护者及其选择的所有依赖。例如，一个JS库的devDependencies、Rust crate的构建工具、CI中的插件，都构成你安装包的信任链的一部分。
  - WHY: 现代软件构建依赖数百个间接包。若任一环节的维护者疏忽（如未审计依赖、使用陈旧工具），你的生产环境可能被植入后门。Thompson的编译器后门思想在开源生态中重现，且更难追踪。
  - ACTION: 1. 审计完整依赖树，包括devDependencies和构建工具。2. 使用SBOM（软件物料清单）工具如Syft、SPDX可视化全链。3. 对关键依赖，定期重建工具链并对比输出。4. 优先选择维护活跃、依赖精简的项目。
  - TWEET: 开源依赖的传递信任：你不仅信任直接依赖，还信任其维护者及其所有选择。一个疏忽的devDependency就可能污染整个链。立即行动：生成SBOM，审计完整树。#安全 #基础设施
