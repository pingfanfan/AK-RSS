# OPMLWatch Digest

Generated at: 2026-02-27 19:12:04 UTC

- [How to Block the ‘Upgrade to Tahoe’ Alerts and System Settings Indicator](https://robservatory.com/block-the-upgrade-to-tahoe-alerts-and-system-settings-indicator/) (`daringfireball.net`)
  - TLDR: 通过macOS设备管理配置文件，可屏蔽Tahoe升级提示90天，不影响其他更新。
  - WHAT: 利用macOS设备管理配置文件中的‘限制主要OS更新’策略，隐藏系统设置中的Tahoe升级选项及红点提示，有效期90天。
  - WHY: 避免误触升级导致开发环境中断，保持当前系统（如Sequoia）稳定，减少干扰，专注工作。
  - ACTION: 1. 创建.mobileconfig文件，添加com.apple.applicationaccess策略，设置forceDelayedSoftwareUpdates为true并设置delayDuration=90。2. 删除其他相关行（如forceReleaseDateLock）以确保仅屏蔽主要更新。3. 安装配置文件到目标Mac。4. 90天后需重新安装。
  - TWEET: macOS开发者注意：通过设备管理配置文件可屏蔽Tahoe升级提醒90天，不影响其他更新。避免误操作，保持Sequoia环境稳定。教程链接：https://robservatory.com/block-the-upgrade-to-tahoe-alerts-and-system-settings-indicator/ #安全 #基础设施
