# OPMLWatch Digest

Generated at: 2026-03-06 10:43:07 UTC

- [I don't know if my job will still exist in ten years](https://seangoedecke.com/will-my-job-still-exist/) (`seangoedecke.com`)
  - TLDR: AI代理将重塑软件工程：资深工程师或转型为监督者，初级岗位首当其冲；未来十年取决于行业对AI能力的判断是保守还是激进。
  - WHAT: 核心问题是科技行业对AI代理能力的预期偏差——低估则缓慢替代，高估则引发人才短缺。
  - WHY: AI正快速改变工程工作本质，初级/中级岗位首当其冲；理解这一偏差能帮助个人提前规划转型，避免职业风险。
  - ACTION: 立即开始学习AI代理协作与监督技能，深化系统设计能力，并拓展至安全、产品等交叉领域，构建“人机协同”的不可替代性。
  - TWEET: 软件工程师的未来：AI代理普及后，写代码减少，监督AI增加。你的岗位会先被影响吗？
- [How I think systemd IP address restrictions on socket units works](https://utcc.utoronto.ca/~cks/space/blog/linux/SystemdSocketIPRestrictions) (`utcc.utoronto.ca/~cks`)
  - TLDR: 在systemd .socket单元设置IPAddressAllow/Deny，仅限制对该socket的入站访问，不影响服务自身出站连接（如DNS查询）。
  - WHAT: systemd的IP地址控制（IPAddressAllow/Deny）通过eBPF实现。当配置在.socket单元时，规则仅应用于该socket的入站流量，服务进程的出站流量不受影响。
  - WHY: 这提供了精细的入站访问控制，避免了在.service单元上设置IP限制时，会同时阻断服务必要出站连接（如API调用、DNS）的问题。实现了“只过滤谁可以连我，不限制我可以连谁”。
  - ACTION: 若需为socket激活服务（如ssh.socket）配置入站IP白名单，直接在.socket单元文件中使用IPAddressAllow=，而非在对应的.service单元中设置。验证时可检查对应cgroup下的eBPF程序（bpftool cgroup list）。
  - TWEET: 在systemd .socket单元设置IP限制，竟只影响入站？是的！这通过eBPF直接附着在socket上实现，服务进程的出站连接（如DNS、API调用）完全自由。这是为socket激活服务设计入站防火墙的优雅方案。验证：bpftool cgroup list | grep sd_fw。
- [.gitlocal](https://nesbitt.io/2026/03/06/gitlocal.html) (`nesbitt.io`)
