# OPMLWatch Digest

Generated at: 2026-03-10 06:44:27 UTC

- [Power glitches can leave computer hardware in weird states](https://utcc.utoronto.ca/~cks/space/blog/tech/PowerGlitchesWeirdHardwareStates) (`utcc.utoronto.ca/~cks`)
  - TLDR: 短暂电力故障可能使服务器、交换机等硬件进入通电但功能异常的状态，需完全断电才能恢复，这暴露了数字系统底层模拟电路的脆弱性。
  - WHAT: 硬件在遭遇电力波动后，可能处于“假死”状态：设备指示灯正常，但网络端口、管理接口（如BMC/IPMI）无响应，常规软重启无效，必须彻底断电再通电。并非所有同型号设备都会中招，存在随机性。
  - WHY: 电力瞬变可能将部分模拟电路（如时钟、电压调节器、信号电平）推至非预期状态，数字逻辑无法检测或自愈。残留“跳蚤电力”（flea power）可能维持异常状态，阻止系统完全重置。
  - ACTION: 1. 关键设备部署UPS并监控电源质量；2. 设计硬件看门狗与带外管理（如IPMI）的强制断电逻辑；3. 制定电力故障后完全断电重启的SOP；4. 测试不同硬件型号对瞬态电源的敏感性。
  - TWEET: 电力瞬变可能让硬件陷入无法自愈的异常状态，即使通电也无效。关键：1. UPS监控 2. 带外管理强制断电 3. 制定断电重启SOP。你的数据中心遇到过吗？ #DevOps #SRE
