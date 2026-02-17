# OPMLWatch Digest

Generated at: 2026-02-17 08:44:34 UTC

- [Understanding the limitation of 'do in new frame/window' in GNU Emacs](https://utcc.utoronto.ca/~cks/space/blog/programming/EmacsNewWhateverLimitation) (`utcc.utoronto.ca/~cks`)
  - TLDR: Emacs 的“新框架/窗口”功能基于“显示缓冲区”模型，而非“运行命令”，因此没有直接的“在新框架执行命令”函数。
  - WHAT: 博客解析了 `C-x 5 5` 前缀命令的工作原理：它钩住后续命令尝试显示缓冲器的操作，在此刻创建新框架并显示该缓冲器，而非先执行命令。
  - WHY: Emacs 核心模型中，框架/窗口是缓冲区的“显示容器”，创建框架必须已知要显示的缓冲区。命令执行可能产生多个缓冲区，无法预设。
  - ACTION: 在 Elisp 中实现“在新框架运行命令”，需先确保命令产生的缓冲区已存在，再使用 `(switch-to-buffer-other-frame (current-buffer))` 或类似函数显示它。
  - TWEET: Emacs 新框架限制解析：框架显示缓冲区，不直接运行命令。`C-x 5 5` 是钩住“显示缓冲区”事件来创建框架。自定义功能时，先让命令生成缓冲区，再用 `switch-to-buffer-other-frame` 显示。适用于 MH-E、Magit 等非文件缓冲区场景。
