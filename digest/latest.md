# OPMLWatch Digest

Generated at: 2026-02-15 19:37:20 UTC

- [Race between primes of the forms 4k + 1 and 4k + 3](https://www.johndcook.com/blog/2026/02/15/chebyshev-bias/) (`johndcook.com`)
  - TLDR: 最大素数纪录多为4k+3形式，因Mersenne素数更易通过Lucas-Lehmer测试发现，且存在“Chebyshev偏差”——有限范围内4k+3素数计数常领先，但极限下两类素数数量相等。
  - WHAT: 博客分析了素数模4分布，解释世界纪录素数（如Mersenne素数）为何多为4k+3，并引入Chebyshev偏差：计数差g(n)会无限振荡，但4k+3常长期领先。
  - WHY: 对开发者至关重要：密码学参数（如Curve25519的4k+1素数）选择需理解分布；Chebyshev偏差警示勿因短期数据误判数论规律。
  - ACTION: 1. 运行文末Python代码计算g(n)，可视化偏差现象；2. 评估密码系统时，明确素数形式对算法（如平方和表示）的影响；3. 关注素数纪录时，区分“易发现”与“数量多”的数学原因。
  - TWEET: 为何最大素数纪录全是4k+3？因Mersenne素数易用Lucas-Lehmer测试发现，非因数量更多。数论中存在“Chebyshev偏差”：有限范围内4k+3素数计数常领先，但极限下与4k+1相等，且会无限振荡。附可视化代码，开发者需注意：密码学选参时，素数形式影响算法特性（如平方和表示），勿被短期数据误导。
