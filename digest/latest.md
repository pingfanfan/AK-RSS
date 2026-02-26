# OPMLWatch Digest

Generated at: 2026-02-26 10:44:26 UTC

- [Git in Postgres](https://nesbitt.io/2026/02/26/git-in-postgres.html) (`nesbitt.io`)
  - TLDR: 本文探讨用Postgres替代Git传统存储，解决包管理器扩展性问题，实现更高效协作。
  - WHAT: 一种将Git对象存储后端替换为Postgres数据库的技术方案，通过两张表实现完整Git数据模型。
  - WHY: 解决Git在大规模仓库中的性能瓶颈（如Homebrew被迫停止全量克隆），利用数据库的ACID、压缩和并发优势。
  - ACTION: 开发者可探索将Postgres后端集成到内部Git服务，或评估现有工具（如git-base）的适用性。
  - TWEET: Git存储瓶颈如何破？用Postgres做后端，两张表搞定Git数据模型。Homebrew案例启示：别再用Git当数据库了。技术前沿探讨。
