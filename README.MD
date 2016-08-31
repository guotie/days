# 计算自然日的工具集

## Tomorrow(tm time.Time) time.Time

返回明天00：00：00的时间

## Yesterday(tm time.Time) time.Time

返回昨天00:00:00的时间

## Today() time.Time

返回今天00:00:00的时间

## DaysBetween(t1, t2 time.Time) (days int)

计算两个时间t1, t2相差的自然天

## DaysAfter(tm time.Time, n int) time.Time

计算距离tm以后n天的自然天的00:00:00的时间

## DaysBefore(tm time.Time, n int) time.Time

计算距离tm以前n天的自然天的00:00:00的时间


# timer

    func StartTimer(n, m int64, fn func(args ...interface{}) error, args ...interface{}) *time.Timer
	
	n, m 为毫秒数