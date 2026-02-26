# OPMLWatch Digest

Generated at: 2026-02-26 07:19:24 UTC

- [With disk caches, you want to be able to attribute hits and misses](https://utcc.utoronto.ca/~cks/space/blog/tech/DiskCacheHitMissWantDetails) (`utcc.utoronto.ca/~cks`)
  - TLDR: 仅看整体缓存命中率会掩盖关键性能差异。必须细分I/O类型（如数据/元数据、同步/异步）才能精准优化，但OS层常缺乏此细节。
  - WHAT: 问题在于：磁盘缓存的不同I/O源（文件数据、各种元数据）命中率可能天差地别，笼统的命中率指标会误导性能分析，让你误以为系统健康。
  - WHY: 因为高命中率可能只针对特定元数据，一旦该命中率下降（如工作集变化），相关性能会急剧恶化。细粒度指标是定位瓶颈、评估改进效果的唯一途径。
  - ACTION: 1. 检查现有监控：是否区分I/O类型（数据/元数据、同步/异步、预读/需求读）？2. 在ZFS等系统中启用细粒度命中率指标。3. 设计新缓存时，将可归因性作为核心埋点需求。
  - TWEET: 别被整体缓存命中率欺骗。元数据访问可能近乎100%命中，但文件I/O可能频繁未命中。性能瓶颈常藏在细分数据里。检查你的指标能否归因I/O类型。
