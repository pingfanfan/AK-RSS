# OPMLWatch Digest

Generated at: 2026-02-20 07:18:09 UTC

- [Parsing hours and minutes into a useful time in basic Python](https://utcc.utoronto.ca/~cks/space/blog/python/HoursMinutesToTime) (`utcc.utoronto.ca/~cks`)
  - TLDR: 解析‘HH:MM’时，time.strptime()默认年份1900，需结合今日日期修正。推荐用datetime.combine()。
  - WHAT: Python的time.strptime('%H:%M')返回struct_time的年份固定为1900，无法直接代表‘今天此时’。这导致计算相对时间戳时出错。
  - WHY: 开发者常需实现‘--at HH:MM’功能（如查看过去某时刻状态）。标准库此行为源于C库历史，但不符合现代应用需求，需手动修正。
  - ACTION: 1. 用datetime.datetime.combine(datetime.date.today(), parsed_time) 2. 转换为本地时间戳：timestamp = int(combined_dt.timestamp()) 3. 注意时区：若需UTC，用datetime.timezone.utc。
  - TWEET: Python时间解析陷阱：strptime('%H:%M')返回1900年！要得到‘今天此时’的时间戳，请用datetime.combine(today, parsed_time)修正。这是标准库的C遗留问题，实用方案在此。
