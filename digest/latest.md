# OPMLWatch Digest

Generated at: 2026-03-04 06:44:22 UTC

- [A taxonomy of text output (from tools that want to be too clever)](https://utcc.utoronto.ca/~cks/space/blog/sysadmin/ProgramTextOutputTaxonomy) (`utcc.utoronto.ca/~cks`)
  - TLDR: 文章将命令行工具的输出方式从优到劣分类，强调简单、无控制字符的输出更利于自动化处理和日志分析，批评了apt-get等工具的进度条动画。
  - WHAT: 作者基于个人经验，列出六种文本输出方式：从理想的纯文本逐行输出，到使用点号、退格、重复行、颜色，直至最差的进度条动画，分析每种对工具链、日志捕获和分析的影响。
  - WHY: 对于关注基础设施和自动化的开发者，输出格式直接影响工具集成、日志解析和故障排查；过于花哨的输出会破坏管道操作、增加分析复杂度，并与非交互式环境（如CI/CD）冲突。
  - ACTION: 优先选择或设计输出简洁的工具，避免颜色和控制字符；在脚本中处理输出时，使用工具（如script、less -r）或预处理（如sed）净化输出；向工具开发者反馈，要求提供关闭“花哨”输出的选项。
  - TWEET: apt-get的进度条动画真的有必要吗？它让日志难以解析，干扰脚本处理。作者将文本输出分级：纯文本最优，动画最差。作为开发者，我们应倡导工具输出对机器友好。点击阅读全文：https://utcc.utoronto.ca/~cks/space/blog/sysadmin/ProgramTextOutputTaxonomy
- [Remove annoying website elements](https://maurycyz.com/projects/fixsite/) (`maurycyz.com`)
  - TLDR: 一段JS代码，一键移除固定/粘性定位元素，恢复被禁用的滚动，告别弹窗与遮挡。
  - WHAT: 一个书签工具（bookmarklet），遍历页面DOM，检测并移除position为fixed/sticky的元素，同时将overflow设为visible以恢复滚动。
  - WHY: 现代网页常充斥侵入式UI（如Cookie横幅、推荐侧边栏），干扰阅读。此工具可快速清理，恢复页面纯净度与可控性。
  - ACTION: 将提供的JS代码保存为浏览器书签，访问任意网页时点击书签即可一键清理。
  - TWEET: 网页太乱？用这个JS书签一键清理固定定位元素和禁用滚动。实用工具，推荐给所有经常上网的开发者。#前端开发
