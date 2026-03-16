# OPMLWatch Digest

Generated at: 2026-03-16 20:43:19 UTC

- [Quoting Guilherme Rambo](https://simonwillison.net/2026/Mar/16/guilherme-rambo/#atom-everything) (`simonwillison.net`)
  - TLDR: MacBook Neo的软件摄像头指示灯运行在Secure Enclave安全隔离区，即使内核被攻破也无法关闭，提供近乎硬件的隐私保护。
  - WHAT: 摄像头指示灯作为一个独立进程，在Secure Enclave（与内核隔离的 privileged 环境）中运行，并直接操作屏幕硬件像素，而非通过操作系统渲染。
  - WHY: 传统软件指示灯可能被恶意内核驱动或漏洞关闭。此设计将控制权置于最高安全等级，确保任何试图访问摄像头的操作都会强制点亮指示灯，从根本上防止隐蔽偷拍。
  - ACTION: 评估你负责系统的关键安全组件（如权限管理、审计日志）是否可借鉴此类“安全隔离区”设计，将核心逻辑与主系统隔离，降低提权攻击面。
  - TWEET: 苹果把摄像头指示灯逻辑塞进Secure Enclave，直接blit到屏幕。这意味着内核级漏洞也无法关闭它。这是把“隐私可视化”做到了架构底层。开发者：你的关键安全逻辑是否也足够隔离？
