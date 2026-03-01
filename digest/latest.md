# OPMLWatch Digest

Generated at: 2026-03-01 07:40:44 UTC

- [Sometimes the simplest version of a text table is printed from a command](https://utcc.utoronto.ca/~cks/space/blog/sysadmin/SimpleTextViaCommands) (`utcc.utoronto.ca/~cks`)
  - TLDR: 停电故障后，用命令行脚本直接查询Prometheus获取宕机主机列表，比在Grafana网页中查看表格更快速、更适合服务器console操作。
  - WHAT: 一个名为`promdownhosts`的脚本，通过curl查询Prometheus API，用jq提取数据，打印出当前宕机的主机名列表。
  - WHY: 在紧急故障排查场景（如停电后），运维人员常在服务器console操作，无浏览器环境。命令行工具能直接、快速获取关键数据，且易于过滤和自动化，效率远高于打开网页。
  - ACTION: 为你常用的监控查询（如Prometheus）编写小型命令行脚本，用curl+jq处理，输出纯文本。将脚本团队共享，放入统一位置，方便紧急时快速调用。
  - TWEET: 紧急运维中，命令行脚本的价值凸显：无需浏览器，直接查询Prometheus并输出宕机列表，速度与灵活性完胜网页仪表盘。你的团队有类似的‘急救’脚本吗？
