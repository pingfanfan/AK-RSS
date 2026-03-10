# OPMLWatch Digest

Generated at: 2026-03-10 16:18:04 UTC

- [Simplifying expressions in SymPy](https://www.johndcook.com/blog/2026/03/10/simplifying-expressions-in-sympy/) (`johndcook.com`)
  - TLDR: SymPy简化sinh(acosh(x))等表达式时，默认返回sqrt(x-1)*sqrt(x+1)而非sqrt(x**2-1)。需通过symbols声明x为非负实数才能获得预期简化结果。
  - WHAT: 本文通过Python SymPy库验证双曲-反双曲函数简化对照表，并揭示sqrt(x-1)*sqrt(x+1)不自动合并为sqrt(x**2-1)的技术细节。
  - WHY: 根本原因在于复数域中，根号乘积性质√(a)√(b)=√(ab)并非普遍成立。例如x=-2时，左边为√(-3)√(-1)=i√3*i=-√3，右边为√(3)=√3，符号相反。
  - ACTION: 在SymPy中，使用symbols('x', real=True, nonnegative=True)明确变量范围，即可让simplify()正确输出sqrt(x**2-1)。
  - TWEET: 用SymPy做数学简化？注意：sqrt(x-1)*sqrt(x+1) ≠ sqrt(x**2-1) 对所有x成立。指定x≥0即可解决。技术细节与Mathematica行为一致。
