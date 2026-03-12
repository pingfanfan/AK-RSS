# OPMLWatch Digest

Generated at: 2026-03-12 12:45:02 UTC

- [Historic Energy Price Cap Data (FOI success!)](https://shkspr.mobi/blog/2026/03/historic-energy-price-cap-data-foi-success/) (`shkspr.mobi`)
  - TLDR: 作者通过信息自由请求从Ofgem获取完整历史电价数据，发现原始数据格式混乱，已用R语言清理并开源分享。
  - WHAT: 这是一份通过官方渠道（FOI）获取的英国各地区历史电价上限数据集，覆盖从引入至今的每千瓦时价格，并附有原始与清理后的版本。
  - WHY: 展示了开发者如何主动获取政府开放数据，并处理现实中的脏数据（日期格式混乱、精度冗余），最终以合规方式（OGL 3.0）重新发布，为能源分析、模型训练提供基础。
  - ACTION: 下载清理后的CSV数据用于分析；参考作者R脚本实践数据清洗；若需其他历史数据，可尝试类似FOI请求。
  - TWEET: 开发者如何获取政府历史数据？案例：向英国能源监管机构Ofgem发送FOI邮件，一个月后收到原始Excel，但数据有13位小数、日期格式不统一。作者用R清洗后发布合规CSV（OGL 3.0）。实践要点：1.善用FOI 2.掌握数据清洗 3.确认开源许可。下载：shkspr.mobi/blog/2026/03/...
- [The Elusive Cost Savings of the Prefabricated Home](https://www.construction-physics.com/p/the-elusive-cost-savings-of-the-prefabricated) (`construction-physics.com`)
