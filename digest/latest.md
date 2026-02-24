# OPMLWatch Digest

Generated at: 2026-02-24 06:46:03 UTC

- [Gnome, GSettings, gconf, and which one you want](https://utcc.utoronto.ca/~cks/space/blog/linux/DconfVsGconfInGnome) (`utcc.utoronto.ca/~cks`)
  - TLDR: 现代GNOME程序使用GSettings/dconf（工具：dconf-editor/gsettings），旧程序使用GConf（工具：gconf-editor/gconftool-2）。选错工具将无法修改配置。
  - WHAT: GNOME有两代配置系统：新的是GSettings（后端为dconf），通过dconf-editor图形工具或gsettings命令行修改；旧的是GConf，通过gconf-editor或gconftool-2修改。例如gnome-terminal用新系统，Thunderbird（截至2024）仍用旧系统。
  - WHY: 混淆工具会导致配置无效，影响开发调试、系统管理和自动化脚本。了解系统演进有助于精准定位配置项，避免在错误的地方浪费时间。
  - ACTION: 1. 判断程序类型：现代GNOME原生程序用GSettings/dconf，未更新程序（如Thunderbird）用GConf。2. 优先使用命令行工具（gsettings/dconf/gconftool-2）进行脚本化操作。3. 修改前备份对应数据库。
  - TWEET: GNOME配置系统演进：GSettings/dconf vs GConf。用错工具=白忙活。现代程序（如gnome-terminal）走新系统，Thunderbird等老程序仍用GConf。命令行工具更高效，修改前记得备份。#开发者 #系统管理
