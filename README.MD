# 计算自然日的工具集

## Tomorrow(tm time.Time) time.Time

返回明天00：00：00的时间

## Yesterday(tm time.Time) time.Time

返回昨天00:00:00的时间

## Today() time.Time

返回今天00:00:00的时间

## Between(t1, t2 time.Time) (days int)

计算两个时间t1, t2相差的自然天

## After(tm time.Time, n int) time.Time

计算距离tm以后n天的自然天的00:00:00的时间

## Before(tm time.Time, n int) time.Time

计算距离tm以前n天的自然天的00:00:00的时间


# timer

    func StartTimer(n, m int64, fn func(args ...interface{}) error, args ...interface{}) *time.Timer
	
	n, m 为毫秒数

    启动定时器
    n为第一次执行距离当前时间的毫秒数
    m为以后每次执行的间隔时间，单位毫秒，如果为0，则定时器仅执行一次
    fn 定时器执行函数


    func StopTimer(tmr *Timer)

    停止定时器
