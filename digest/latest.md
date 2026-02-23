# OPMLWatch Digest

Generated at: 2026-02-23 18:16:59 UTC

- [Customizing the ways the dialog manager dismisses itself: Detecting the ESC key, second (failed) attempt](https://devblogs.microsoft.com/oldnewthing/20260223-00/?p=112080) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: 使用GetKeyState检测ESC键状态无法可靠区分IDCANCEL是由ESC键还是关闭按钮触发，因为按键状态和事件生成存在竞态条件，可能导致误判。
  - WHAT: 问题在于如何准确识别IDCANCEL事件的来源：是用户按下ESC键，还是点击了窗口关闭按钮？开发者尝试用GetKeyState检查ESC键在事件发生时的状态，但方法有缺陷。
  - WHY: 失败原因：GetKeyState返回的是按键在调用时的状态，而非事件生成时的状态。如果用户同时按住ESC并点击关闭按钮，第一个IDCANCEL来自ESC（应忽略），但第二个来自关闭按钮时ESC仍被按住，GetKeyState会误判为来自ESC，导致错误忽略。这类似于‘两门问题’：后门开着时有人从前门进入，逻辑会错误抑制动作。
  - ACTION: 不要依赖GetKeyState或GetAsyncKeyState来区分ESC和关闭按钮。应使用更可靠的方法，如检查WM_COMMAND消息的lParam（对于按钮点击）或利用对话框消息处理中的特定标志。参考后续文章或Windows文档中的最佳实践。
  - TWEET: Windows对话框开发中，用GetKeyState检测ESC键来区分IDCANCEL来源？第二次尝试依然失败。竞态条件导致误判：ESC按住时点关闭按钮，第二个IDCANCEL会被错误忽略。正确方法应检查消息参数而非按键状态。
