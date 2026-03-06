# OPMLWatch Digest

Generated at: 2026-03-06 18:06:54 UTC

- [When Read­Directory­ChangesW reports that a deletion occurred, how can I learn more about the deleted thing?](https://devblogs.microsoft.com/oldnewthing/20260306-00/?p=112116) (`devblogs.microsoft.com/oldnewthing`)
  - TLDR: ReadDirectoryChangesW的删除通知只提供文件名，无法获取文件类型/大小。解决方案：改用ReadDirectoryChangesExW并请求扩展信息，删除通知中将包含文件属性。
  - WHAT: 当ReadDirectoryChangesW报告FILE_ACTION_REMOVED时，由于项目已被删除，无法通过GetFileAttributesEx查询其属性（如是否为目录、文件大小），导致监控缓存无法准确更新。
  - WHY: ReadDirectoryChangesW设计用于检测目录列表中的变更，依赖先通过FindFirstFile/FindNextFile构建缓存，再通过增量更新维护。但若项目被快速添加并删除，缓存中无记录，删除时无法获取信息。
  - ACTION: 若需在删除通知中获取被删项详情，应改用ReadDirectoryChangesExW，并在dwNotifyFilter中设置ReadDirectoryNotifyExtendedInformation。这样FILE_ACTION_REMOVED通知将包含FILE_NOTIFY_EXTENDED_INFORMATION结构，其中有FileAttributes和FileSize。
  - TWEET: 监控目录删除时，ReadDirectoryChangesW通知里只有文件名？想获取被删文件的类型和大小？升级到ReadDirectoryChangesExW并开启扩展信息，删除事件直接包含属性。
