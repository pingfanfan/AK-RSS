# OPMLWatch Digest

Generated at: 2026-02-14 22:39:16 UTC

- [Finding a square root of -1 mod p](https://www.johndcook.com/blog/2026/02/14/square-root-minus-1-mod-p/) (`johndcook.com`)
  - TLDR: 当奇素数 p ≡ 1 mod 4 时，-1 在模 p 下有平方根，可通过先找非二次剩余 c，再计算 c^((p-1)/4) mod p 得到。
  - WHAT: 具体步骤：1. 验证 p % 4 == 1；2. 随机或按策略找到模 p 下的非二次剩余 c（如尝试小整数）；3. 计算 k = (p-1)//4，求 x = pow(c, k, p)。
  - WHY: 由欧拉准则，c 是非二次剩余意味着 c^((p-1)/2) ≡ -1 mod p。结合 p=4k+1，则 (c^k)^2 = c^(2k) = c^((p-1)/2) ≡ -1 mod p，故 x=c^k 即为所求。
  - ACTION: 直接运行以下 Python 代码验证示例：`p = 2**255 - 19; c = 2; k = (p-1)//4; x = pow(c, k, p); print((x*x + 1) % p)` 结果应为 0。
  - TWEET: 如何在模素数下求 -1 的平方根？条件：p ≡ 1 mod 4。方法：先找非二次剩余 c（如 2），计算 x = c^((p-1)/4) mod p。例如 p=2^255-19，x=pow(2, (p-1)//4, p) 即满足 x² ≡ -1。代码验证秒出结果！#密码学 #Python #数论
- [Microsoft Game Pass Ultimate Billing Fraud](https://jayd.ml/2026/02/14/microsoft-game-pass-fraud.html) (`jayd.ml`)
  - TLDR: 微软被指在用户明确关闭自动续费后，仍于三年后偷偷开启并扣款，暴露订阅系统安全与黑暗模式风险。
  - WHAT: 用户购买Game Pass Ultimate时彻底关闭自动续费并启用虚拟卡，但三年后订阅被自动续费扣款，质疑微软欺诈。
  - WHY: 可能因系统漏洞、故意设计（黑暗模式）或支付验证失效，反映订阅管理产品在用户控制与透明度上的缺陷。
  - ACTION: 1. 定期复查订阅状态；2. 始终使用一次性虚拟卡；3. 关闭自动续费后截图存证；4. 向消协或网信办举报。
  - TWEET: 微软Game Pass Ultimate被曝用户关闭自动续费后仍被偷偷开启扣款！安全警示：订阅服务需定期检查，务必用虚拟卡并截图保存设置。这不仅是漏洞，更是产品黑暗模式。建议向12315举报。 #订阅欺诈 #产品安全 #微软
