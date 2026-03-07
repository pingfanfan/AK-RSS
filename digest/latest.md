# OPMLWatch Digest

Generated at: 2026-03-07 07:07:34 UTC

- [Your terminal program has to be where xterm's ziconbeep feature is handled](https://utcc.utoronto.ca/~cks/space/blog/unix/XTermHasToDoZiconbeep) (`utcc.utoronto.ca/~cks`)
  - TLDR: 窗口管理器无法可靠替代 xterm 的 ziconbeep 功能，因为无法准确检测窗口是否有新输出，且不理解变化的语义意义。该逻辑必须由终端程序自身实现。
  - WHAT: ziconbeep 是 xterm 在窗口被最小化且接收到新输出时，通过改变标题或发出提示来吸引用户注意的功能。有观点认为此功能应移至更通用的窗口管理器层实现。
  - WHY: 1. 输出检测难：X/Wayland 协议不提供窗口渲染状态的可靠可见性，现代 GUI 常持续渲染（即使最小化），窗口管理器无法区分‘有意义的新输出’与‘常规重绘’。2. 语义判断需程序自身：只有程序知道哪些变化值得提示（如邮件客户端只对特定邮件提示）。
  - ACTION: 开发者若需实现类似 ziconbeep 的‘最小化时新输出提示’功能，应在应用程序层（如终端模拟器）主动处理：监听自身输出流，结合窗口状态（是否最小化）和业务逻辑（如过滤无关日志）来触发提示。
  - TWEET: 窗口管理器能响应窗口标题变化并播放声音，但无法判断‘变化是否由新输出引起’。xterm 的 ziconbeep 依赖终端自身对输出和最小化状态的感知。将此类逻辑上移给 WM 在技术上不可行。讨论：https://utcc.utoronto.ca/~cks/space/blog/unix/XTermHasToDoZiconbeep
