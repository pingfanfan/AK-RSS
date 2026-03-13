# OPMLWatch Digest

Generated at: 2026-03-13 15:43:54 UTC

- [Typesetting sheet music with AI](https://www.johndcook.com/blog/2026/03/13/typesetting-sheet-music-with-ai/) (`johndcook.com`)
  - TLDR: 测试AI将乐谱图片转为Lilypond代码：生成的乐谱完全错误，但AI成功识别了曲作者、风格甚至标题，凸显其模式识别能力远超精确代码生成。
  - WHAT: 一项实证测试：评估大语言模型在罕见领域（Lilypond音乐排版语言）的代码生成能力。
  - WHY: Lilypond公开代码极少，测试AI在低资源领域的表现；结果揭示AI依赖统计模式识别，而非对符号的精确理解，这对安全关键型代码生成有警示意义。
  - ACTION: 开发者：勿依赖AI生成小众语言精确代码（如硬件描述、形式化验证），但可尝试用于灵感或识别；务必人工审查输出，尤其在安全关键场景。
  - TWEET: AI 排版乐谱测试：输入乐谱图，输出代码完全错误，但AI猜出了曲作者和风格。这暴露了其在低资源语言中的根本局限——模式匹配强，精确生成弱。安全关键代码？务必人工复核。
