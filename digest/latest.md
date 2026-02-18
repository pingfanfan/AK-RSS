# OPMLWatch Digest

Generated at: 2026-02-18 21:09:50 UTC

- [Frigate with Hailo for object detection on a Raspberry Pi](https://www.jeffgeerling.com/blog/2026/frigate-with-hailo-for-object-detection-on-a-raspberry-pi/) (`jeffgeerling.com`)
  - TLDR: Hailo AI加速器（HAT+或M.2）为树莓派Frigate NVR提供了比Coral TPU更强大、更灵活的硬件目标检测方案。
  - WHAT: 博客介绍了将树莓派5的官方AI HAT+（内置Hailo-8/8L）或通用M.2 Hailo卡，用于替代现有Coral TPU，在Frigate中实现低功耗、高性能的实时对象检测。
  - WHY: Hailo性能显著优于Coral，支持更复杂的模型，且M.2接口使其能用于更多非树莓派设备，解决了Coral生态封闭和性能瓶颈问题。
  - ACTION: 1. 评估Hailo-8L（性价比）或Hailo-8（性能）是否满足你的检测精度/延迟需求；2. 确认你的Frigate版本已支持Hailo插件；3. 根据你的主板（Pi5或其它SBC）购买对应Hailo硬件。
  - TWEET: 树莓派Frigate安防新选择：用Hailo AI加速器（HAT+/M.2）替代Coral TPU！性能更强，模型支持更广，且M.2版本通用性高。升级指南：评估需求→确认Frigate兼容→选购硬件。低功耗高效能AI推理，为你的智能家居安防注入新动力。#Frigate #Hailo #树莓派 #AIoT #智能安防
