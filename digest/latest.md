# OPMLWatch Digest

Generated at: 2026-02-14 21:37:41 UTC

- [Justifying text-wrap: pretty](https://matklad.github.io/2026/02/14/justifying-text-wrap-pretty.html) (`matklad.github.io`)
  - TLDR: Safari 2025年新增CSS属性`text-wrap: pretty`，使用动态规划优化网页文本换行，但当前与`text-align: justify`连用会导致词间距异常。
  - WHAT: `text-wrap: pretty`是一个CSS属性，旨在通过智能选择换行点，使文本行的长度更均衡，提升排版美观度。
  - WHY: 传统浏览器使用简单贪婪算法导致排版不佳；该属性将Knuth-Plass算法适配到响应式Web环境，是排版技术的重大进步。
  - ACTION: 立即在Safari中为段落应用`text-wrap: pretty`测试效果；若同时使用`text-align: justify`，需仔细检查并调整词间距，或考虑仅使用`pretty`。
  - TWEET: Safari 已支持 `text-wrap: pretty`！它用动态规划让网页文本换行更均衡，告别丑陋排版。但注意：与 `text-align: justify` 连用时词间距可能爆炸。开发者速测，并关注其他浏览器跟进情况。 #前端 #CSS
