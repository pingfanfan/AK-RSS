# OPMLWatch Digest

Generated at: 2026-03-02 21:42:31 UTC

- [I built a pint-sized Macintosh](https://www.jeffgeerling.com/blog/2026/pint-sized-macintosh-pico-micro-mac/) (`jeffgeerling.com`)
  - TLDR: 使用Raspberry Pi Pico复刻经典Macintosh，支持VGA显示与USB键鼠，内存达208KB，比原版多63%。
  - WHAT: 这是一个基于RP2040的复古计算项目，运行System 5.3固件，通过VGA输出显示，并兼容标准USB外设。
  - WHY: 它提供了低成本体验经典Mac系统的途径，在极简硬件上实现功能复刻，展示了嵌入式开发的创造力，适合硬件爱好者和教育场景。
  - ACTION: 访问GitHub获取源码，购买Raspberry Pi Pico，按指南组装并体验复古计算。
  - TWEET: 袖珍Macintosh诞生！基于Pico，运行System 5.3，外设全兼容。硬件限制下的创新案例，开发者快来尝试。#嵌入式开发
- [ChangeTheHeaders](https://underpassapp.com/news/2025/3/4.html) (`daringfireball.net`)
  - TLDR: Safari拖拽图片常获WebP格式，因Accept请求头含WebP且质量值相同，服务器自主选择。使用ChangeTheHeaders扩展可修改请求头，强制返回PNG/JPEG以确保兼容性。
  - WHAT: 问题本质是HTTP Accept请求头中image/webp与image/png的质量值均为1，服务器根据自身逻辑选择返回WebP，即使URL后缀为.png。URL扩展名在动态服务中无实际约束力。
  - WHY: 开发者需控制输出格式以保证跨平台兼容性。意外获得WebP可能导致发布或托管时兼容性问题，修改请求头是直接有效的解决方案。
  - ACTION: 安装Jeff Johnson的ChangeTheHeaders Safari扩展，在设置中自定义Accept请求头，移除或降低image/webp的优先级，强制服务器返回PNG或JPEG格式。
  - TWEET: Safari拖拽图片得WebP的坑，原来是Accept头里WebP和PNG优先级相同。ChangeTheHeaders扩展能修改请求头，强制返回所需格式，推荐开发者试试。
