# OPMLWatch Digest

Generated at: 2026-03-01 18:42:44 UTC

- [Shell variable ~-](https://www.johndcook.com/blog/2026/03/01/tilde-dash/) (`johndcook.com`)
  - TLDR: `~-`是`$OLDPWD`的快捷方式，可快速引用上一个工作目录，简化目录间文件比较等操作。
  - WHAT: `~-`是Bash shell的内置变量快捷方式，等价于`$OLDPWD`，代表上一次执行`cd`前的目录路径。
  - WHY: 提升命令行效率：在频繁切换的两个目录间操作时，无需重复输入完整路径或记忆`$OLDPWD`，直接使用`~-`即可快速定位，尤其适用于文件对比、脚本编写等场景。
  - ACTION: 立即在终端尝试：1. 切换目录后，用`ls ~-`查看上一个目录内容；2. 使用`diff file.txt ~-/file.txt`对比同名文件；3. 查阅`man bash`搜索`Tilde Expansion`探索更多类似快捷方式。
  - TWEET: 你知道吗？Bash里有个隐藏变量`~-`，等于`$OLDPWD`。在来回切换的两个目录间对比文件？`diff file ~-/file` 一键搞定。功能老但实用，你们平时还用过哪些冷门但好用的Shell技巧？求分享！
- [HN Skins 0.2.0](https://susam.net/code/news/hnskins/0.2.0.html) (`susam.net`)
