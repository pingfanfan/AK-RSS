# OPMLWatch Digest

Generated at: 2026-02-22 06:44:25 UTC

- [The importance of limiting syndication feed requests in some way](https://utcc.utoronto.ca/~cks/space/blog/web/FeedLimitingImportance) (`utcc.utoronto.ca/~cks`)
  - TLDR: 通过HTTP 304和速率限制，每天避免约3.5GB冗余流量，降低服务器负载，是基础设施优化的关键实践。
  - WHAT: 订阅源抓取器应强制使用HTTP条件请求（如ETag/If-Modified-Since）并遵守速率限制（HTTP 429），避免获取未变更的完整XML/JSON数据。
  - WHY: 实测数据：若无限制，每日额外产生3.5GB流量；41%请求因过于频繁被429拒绝（最活跃IP平均每两分钟一次），严重消耗服务器资源。
  - ACTION: 1. 在抓取器中实现ETag/Last-Modified缓存与If-Modified-Since请求头；2. 收到HTTP 429后实施指数退避重试；3. 监控并限制单IP请求频率（如>1次/分钟）。
  - TWEET: 数据：某博客订阅源，7492次200响应（1.26GB），9419次304，11941次429。若无304和限速，将多产生3.5GB流量。限流不仅是礼貌，更是基础设施必需。 #WebDev #SRE
