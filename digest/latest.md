# OPMLWatch Digest

Generated at: 2026-03-13 12:44:58 UTC

- [An odd font rendering bug in Firefox and Safari](https://shkspr.mobi/blog/2026/03/an-odd-font-rendering-bug-in-firefox-and-safari/) (`shkspr.mobi`)
  - TLDR: Firefox和Safari在特定CSS设置下，对使用组合字符（如ọ́）的姓名渲染异常，会错误回退到默认字体，而Chrome正常。
  - WHAT: 问题源于字体文件缺失预组合字符（如U+1EB9 ẹ）。当CSS设置font-weight: normal时，Firefox/Safari无法正确合成组合字符（e + ◌́），导致部分字形（如加粗的带点e）回退到系统字体。Chrome的合成引擎更稳健。
  - WHY: 这直接影响国际化网站的可访问性与品牌一致性。姓名是核心身份标识，渲染错误可能冒犯用户、损害专业形象，且问题隐蔽，仅影响特定浏览器组合。
  - ACTION: 1. 审计Web字体，确保覆盖所有需要的组合字符（U+1EB9, U+1ECD等）。2. 测试font-weight对组合字符的影响，必要时为特定元素指定font-weight: bold。3. 使用font-display: swap避免不可见文本，但注意可能加剧回退。4. 考虑使用polyfill（如fonttools子集化）或提供备选字体栈。
  - TWEET: 发现一个隐蔽的跨浏览器字体问题：当使用Helvetica Now等字体且CSS为font-weight: normal时，Firefox/Safari无法正确渲染带组合点的字符（如Ronkẹ），部分字母会回退到系统字体。Chrome无此问题。这暴露了字体子集化与合成引擎的兼容性陷阱。你的网站是否测试过组合字符？#Web开发 #字体 #可访问性
