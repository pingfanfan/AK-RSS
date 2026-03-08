# OPMLWatch Digest

Generated at: 2026-03-08 06:08:21 UTC

- [Restricting IP address access to specific ports in eBPF: a sketch](https://utcc.utoronto.ca/~cks/space/blog/linux/EBPFPerPortIPRestrictions) (`utcc.utoronto.ca/~cks`)
  - TLDR: 通过扩展LPM map的键以包含端口号，实现一个简单的eBPF程序，从而在cgroup层面进行跨端口的IP访问控制，满足systemd等工具动态生成程序的需求。
  - WHAT: 本文提出了一种在eBPF cgroup socket SKB程序中实现基于端口的IP访问限制的技术方案。核心是利用LPM (Longest Prefix Match) map的键结构灵活性，将16位端口号置于IP地址之前，从而在单个、简单的eBPF程序中同时匹配源/目标IP和端口。
  - WHY: 传统方法（如systemd的IPAddressAllow/Deny）受限于socket单元，难以跨端口统一策略。而直接编写复杂eBPF程序进行多字段匹配，不利于systemd等工具动态生成。该方案通过数据平面（map）的扩展，保持了控制平面（eBPF程序）的简洁性与可动态生成性，同时实现了精细的跨端口过滤。
  - ACTION: 开发者可参考此设计思路，在需要动态网络策略的场景（如容器运行时、服务网格）中实现类似功能。建议结合现有eBPF工具链（如libbpf、bpftrace）进行原型验证，并关注内核对eBPF map类型的支持演进。
  - TWEET: 在eBPF里做端口级IP限制，不想写复杂程序？试试把端口号塞进LPM map的键里。这样程序逻辑极简，map负责复杂匹配，完美适配systemd这类需要动态生成ebpf程序的工具。技术细节：https://utcc.utoronto.ca/~cks/space/blog/linux/EBPFPerPortIPRestrictions #基础设施 #开发者
