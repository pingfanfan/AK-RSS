# OPMLWatch Digest

Generated at: 2026-03-06 16:44:16 UTC

- [The MacBook Neo’s Price, Looking to the Past and Future](https://x.com/ethan_is_online/status/2029331836137291941?s=42) (`daringfireball.net`)
- [How to Host your Own Email Server](https://blog.miguelgrinberg.com/post/how-to-host-your-own-email-server) (`miguelgrinberg.com`)
  - TLDR: 作者实践自建邮件服务器，分享从零配置到邮件被Gmail等主流服务商接受的完整步骤，涵盖SPF/DKIM/DMARC DNS设置、SMTP服务器与SSL/TLS配置，并提供接收邮件的简单方案。
  - WHAT: 本文具体指导如何配置邮件发送（Postfix SMTP、DNS记录、SSL证书）及接收（使用域名别名转发至Gmail等），解决发信信誉与反垃圾邮件机制问题。
  - WHY: 避免依赖Mailgun/SendGrid等大厂服务，降低长期成本与隐私风险；证明自建虽需技术投入但完全可行，且能掌握基础设施控制权。
  - ACTION: 开发者可参考本文在测试域名上实践，重点配置DNS记录与SMTP认证，使用邮件测试工具验证 deliverability，再逐步迁移至生产环境。
  - TWEET: 自建邮件服务器真的难吗？看这位开发者如何从零配置，让邮件通过Gmail等大厂反垃圾检查。核心是DNS记录与SMTP安全配置，接收邮件竟可用别名转发这么简单。技术自主，拒绝绑架。
