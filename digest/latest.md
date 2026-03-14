# OPMLWatch Digest

Generated at: 2026-03-14 06:10:02 UTC

- [Tim Cook: ‘50 Years of Thinking Different’](https://www.apple.com/50-years-of-thinking-different/) (`daringfireball.net`)
  - TLDR: 库克50周年感谢信真诚但平庸，评论指其缺乏‘Think Different’级标语，苹果品牌叙事需新突破。
  - WHAT: 本文评论苹果CEO库克为50周年撰写的公开信，认为其语言简洁真诚却平淡无奇，并对比30年前‘Think Different’广告的永恒影响力，指出当前苹果营销缺乏令人铭记的宣言。
  - WHY: 库克时代苹果聚焦运营效率与股东价值，而非颠覆性品牌叙事；产品迭代虽稳，但缺乏‘改变世界’的宏大故事，导致营销口号难以流传。
  - ACTION: 开发者应关注苹果技术栈的实际演进（如AI、隐私架构），而非仅依赖品牌光环；在自身产品中注入‘非同凡想’精神，用技术解决真实问题。
  - TWEET: 苹果50周年，库克信函真诚却平庸。对比‘Think Different’的永恒，当前营销无记忆点。对开发者启示：少看口号，多看Swift、AI框架、隐私架构的实际演进。
- [Big tech engineers need big egos](https://seangoedecke.com/big-tech-needs-big-egos/) (`seangoedecke.com`)
- [Last-Run Syndication](https://feed.tedium.co/link/15204/17299183/television-first-run-syndication-decline) (`tedium.co`)
- [How to use the Qwen 3.5 LLMs to OCR documents](https://martinalderson.com/posts/how-to-use-qwen-3-5-to-ocr-documents/?utm_source=rss&utm_medium=rss&utm_campaign=feed) (`martinalderson.com`)
  - TLDR: 使用阿里开源的 Qwen 3.5 多模态模型替代 Gemini 进行文档 OCR。实验表明 Qwen3.5-9B 在效果与成本间取得最佳平衡，能有效处理扫描质量差的 PDF，且保护数据隐私。
  - WHAT: 利用 Qwen 3.5 系列的视觉理解能力，将 PDF 每一页转换为图像后输入模型，直接提取文字内容。核心步骤：1. 用 PyMuPDF 快速提取 PDF 页面为图片；2. 选择合适尺寸的 Qwen3.5 模型（推荐 9B）进行推理。
  - WHY: 1. 成本：Gemini Flash 价格持续上涨，处理大量文档时不经济；2. 隐私：敏感文档无需上传至第三方云服务；3. 性能：Qwen3.5 是首个将多模态能力下放到 0.8B/2B 参数级别的开源模型，9B 版本在 OCR 任务上表现稳定，小模型易偏离任务。
  - ACTION: 1. 安装 PyMuPDF 库；2. 下载 Qwen3.5-9B (Q4_K_M 量化版约 5.68GB) 模型；3. 编写脚本：用 PyMuPDF 提取 PDF 页面为图像，调用模型 API 或本地推理接口，传入图像并提取文本；4. 对比不同尺寸模型（0.8B/2B/4B/9B）在自身文档集上的准确率。
  - TWEET: 还在为 Gemini OCR 账单头疼？试试阿里开源的 Qwen 3.5。实测 9B 模型处理扫描 PDF 效果惊艳，关键数据自己掌控。技术栈：PyMuPDF 提取图片 + Qwen 多模态推理。成本与隐私的双赢。👇
