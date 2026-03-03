# OPMLWatch Digest

Generated at: 2026-03-03 20:42:55 UTC

- [Just for fun: A survey of write protect notches on floppy disks and other media](https://devblogs.microsoft.com/oldnewthing/20260303-00/?p=112104) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 文章对比了8英寸、5¼英寸、3½英寸软盘及Iomega Bernoulli Box的写保护机制：8英寸打孔即保护，其余多为‘写启用’设计（贴纸/滑动门），Bernoulli Box用带符号滑块。揭示硬件设计中‘保护’与‘启用’的语义反转。
  - WHAT: 详细描述各介质写保护位置与操作：8英寸软盘垂直插入，顶部缺口打孔即写保护；5¼英寸水平放置，右侧缺口贴纸覆盖即保护；3½英寸右上角孔，滑动门开则写启用；Bernoulli Box底部滑块带⊘符号，滑至该侧写保护；磁带需同时移除螺丝和滑块。
  - WHY: 对开发者启示：1. 硬件设计需考虑用户直觉与防误操作（如3.5英寸滑动门）；2. 安全机制常通过物理状态实现，影响现代存储设备设计；3. 历史演进显示标准化与用户体验的平衡。
  - ACTION: 建议：1. 审视当前产品硬件安全设计是否清晰防误；2. 研究历史介质设计对现代USB写保护开关、SSR安全擦除的启发；3. 在文档中明确术语，避免‘写保护’歧义。
  - TWEET: 软盘写保护机制大揭秘：8英寸打孔保护，5¼英寸贴纸覆盖，3½英寸滑动门，Bernoulli Box符号滑块。看似‘保护’实为‘启用’？硬件设计思维值得开发者深思。
