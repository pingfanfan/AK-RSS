# OPMLWatch Digest

Generated at: 2026-03-18 16:20:22 UTC

- [Windows stack limit checking retrospective: Alpha AXP](https://devblogs.microsoft.com/oldnewthing/20260318-00/?p=112146) (`devblogs.microsoft.com/oldnewthing`)
- [Wayland has good reasons to put the window manager in the display server](https://utcc.utoronto.ca/~cks/space/blog/unix/WaylandAndBuiltinWindowManagers) (`utcc.utoronto.ca/~cks`)
  - TLDR: Wayland将窗口管理器内置，以避免X11式分离导致的输入事件延迟与竞争，确保实时窗口管理。
  - WHAT: Wayland架构将窗口管理器与显示服务器合并，而X11分离两者。这种设计解决了事件处理与窗口管理深度耦合的问题。
  - WHY: 输入事件（如鼠标移动、点击、打字）的流向由窗口管理器决定。若分离，窗口管理器响应慢会导致事件错位（如打字到错误窗口）。内置使其能无延迟拦截和修改所有事件。
  - ACTION: 开发者应利用Wayland的集成架构实现细粒度输入控制（如应用特定快捷键），并评估现有工具链的兼容性。
  - TWEET: Wayland将窗口管理器内置在显示服务器中，核心原因是输入事件与窗口管理深度耦合。分离会导致同步延迟，使打字或点击可能被错误路由。内置实现无延迟拦截，提升响应性与安全性。
