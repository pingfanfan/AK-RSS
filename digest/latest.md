# OPMLWatch Digest

Generated at: 2026-03-05 20:43:12 UTC

- [The mystery of the posted message that was dispatched before reaching the main message loop](https://devblogs.microsoft.com/oldnewthing/20260305-00/?p=112114) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 发布的消息可能在主消息循环启动前就被处理，通常因准备工作中的隐式消息泵（如DDE/COM调用）触发。解决方案：推迟PostMessage，确保所有初始化完成后再发送。
  - WHAT: 现象：窗口创建后立即Post消息，但消息处理器在准备工作完成前执行。原因：准备工作中的某些操作（如跨进程COM或DDE）内部调用消息泵，导致消息被提前分派。系统无“主循环”概念，任何GetMessage/DispatchMessage调用都可能处理消息。
  - WHY: 客户误解：以为PostMessage的消息必须等主消息循环。实际上消息队列是全局的，任何消息泵（包括初始化代码中的隐式泵）都会处理消息。根本原因：准备工作触发了隐式消息循环，提前消耗了队列中的消息。
  - ACTION: 1. 在消息处理函数设断点，检查调用栈，确认哪个操作触发了消息泵。2. 调整代码顺序，确保所有准备工作（尤其是可能触发消息泵的操作）完成后再Post消息。3. 考虑用SendMessage替代（注意阻塞风险）。
  - TWEET: Win32开发者注意：你以为PostMessage会等主循环？错！任何消息泵（包括COM/DDE的隐式泵）都可能提前处理。调试：在WM_XXX断点看调用栈。修复：推迟Post直到初始化完成。 #Win32 #Debugging
