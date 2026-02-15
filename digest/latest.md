# OPMLWatch Digest

Generated at: 2026-02-15 10:40:57 UTC

- [Separating Download from Install in Docker Builds](https://nesbitt.io/2026/02/15/separating-download-from-install-in-docker-builds.html) (`nesbitt.io`)
  - TLDR: 在Dockerfile中分离依赖下载与安装步骤，能显著提升构建缓存命中率，减少冗余下载，节省时间与公共带宽。
  - WHAT: 将包管理器的“下载”（仅根据lockfile获取依赖到本地缓存）与“安装”（从本地缓存离线安装）拆分为两个独立的RUN指令，并调整COPY顺序，使仅lockfile变更时能复用下载层。
  - WHY: 对开发者可加速本地构建迭代；对社区能减少对crates.io等依赖捐赠带宽服务的无效请求，降低其运营成本，体现技术社会责任。
  - ACTION: Go项目直接采用`go mod download`模式；Rust/Python/Ruby项目可尝试`cargo fetch`/`pip download`等命令模拟分离，并关注其官方离线安装支持；在团队内部分享此模式并审查现有Dockerfile。
  - TWEET: Docker构建加速新思路：分离“下载”与“安装”层！仅COPY lockfile后RUN下载，再COPY源码后离线安装，让缓存更精准。Go社区已实践多年，大幅减少冗余请求。其他语言生态（Rust/Python/Ruby）也值得跟进，既提速本地构建，又减轻crates.io等公共仓库的带宽负担。#Docker #DevOps #Go
