# 计算自然日的工具集

# func Tomorrow(tm time.Time) time.Time

返回明天00：00：00的时间

# func Yesterday(tm time.Time) time.Time

返回昨天00:00:00的时间

# func Today() time.Time

返回今天00:00:00的时间

# func DaysBetween(t1, t2 time.Time) (days int)

计算两个时间t1, t2相差的自然天

# func DaysAfter(tm time.Time, n int) time.Time

计算距离tm以后n天的自然天的00:00:00的时间

# func DaysBefore(tm time.Time, n int) time.Time

计算距离tm以前n天的自然天的00:00:00的时间
