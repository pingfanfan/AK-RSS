# OPMLWatch Digest

Generated at: 2026-02-24 16:48:56 UTC

- [Times New Resistance](https://www.abbyhaddican.com/times-new-resistance) (`daringfireball.net`)
  - TLDR: 一款伪装成Times New Roman的字体，通过连字技术自动将'ICE'替换为'the Goon Squad'等敏感词，展示字体层的文本操纵风险。
  - WHAT: 它利用字体连字特性，在渲染时静默替换特定字符串。安装后系统显示为'Times New Roman'（多一个空格），但实际功能不同。
  - WHY: 作为基础设施，字体可被用于无声的抗议或审查，绕过应用层过滤。开发者需警惕字体供应链攻击与文本安全。
  - ACTION: 1. 审计生产环境字体来源；2. 将字体纳入安全威胁模型；3. 探索字体在UX中的创新与风险平衡。
  - TWEET: 字体也能政治表达？Times New Resistance用连字改写文本。技术上巧妙，但想想恶意字体篡改关键指令？#字体技术 #安全讨论
