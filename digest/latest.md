# OPMLWatch Digest

Generated at: 2026-02-24 15:16:05 UTC

- [Customizing the ways the dialog manager dismisses itself: Isolating the Close pathway](https://devblogs.microsoft.com/oldnewthing/20260224-00/?p=112082) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 通过处理WM_COMMAND/IDCANCEL消息，开发者可精确控制Windows对话框的关闭行为，例如禁用ESC键或关闭按钮，防止误操作。
  - WHAT: 本文详解Windows对话框管理器处理关闭请求的流程，揭示所有关闭操作最终通过WM_COMMAND/IDCANCEL消息路由，为定制关闭逻辑提供统一入口。
  - WHY: 理解此机制能避免UI安全风险（如未保存数据丢失），并实现精细的交互控制，提升应用健壮性和用户体验。
  - ACTION: 在对话框过程中添加WM_COMMAND处理程序，检查IDCANCEL；根据业务逻辑决定调用EndDialog或忽略，并禁用相关控件以视觉提示。
  - TWEET: 对话框关闭不再失控：利用WM_COMMAND/IDCANCEL拦截，自定义Windows应用关闭逻辑。安全、灵活，提升用户体验。#编程 #Windows
