# OPMLWatch Digest

Generated at: 2026-02-15 18:43:08 UTC

- [Gwtar: a static efficient single-file HTML format](https://simonwillison.net/2026/Feb/15/gwtar/#atom-everything) (`simonwillison.net`)
  - TLDR: Gwtar 是一种将网页所有资源打包为单个 HTML 文件的新格式，通过浏览器 API 实现资源的按需懒加载，避免因文件过大导致加载缓慢。
  - WHAT: 它是一个自解压的 HTML 归档格式，文件内嵌 tar 压缩数据。页面加载时先用 `window.stop()` 中断，再通过 `PerformanceObserver` 拦截资源请求，从内联 tar 中按需提取并注入资源。
  - WHY: 解决了传统“单文件 HTML”因内联所有资源（如图片、视频）而体积巨大、难以直接浏览的问题，在保持极致便携（单文件）的同时，实现了类似分块加载的高效体验。
  - ACTION: 开发者可尝试用 Gwtar 打包大型静态站点或离线文档，测试其在不同浏览器中的加载性能；也可研究其源码，借鉴其“中断+拦截+范围请求”的思路用于自定义归档工具。
  - TWEET: 【技术突破】Gwtar 让“单文件网页”不再笨重！它巧妙利用 `window.stop()` 中断初始加载，将全部资源存入内联 tar，再通过 `PerformanceObserver` 拦截请求，按需从 tar 中提取并注入。完美平衡便携性与性能，适合离线文档、大型静态站点归档。前端优化新思路！#WebDev #前端优化 #归档格式
