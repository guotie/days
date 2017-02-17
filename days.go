package days

import (
	"time"
)

// ThisMonth 与tm相同的本月开始时间
func ThisMonth(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, tm.Location())
}

// MonthBefore n个月前的那个月的开始时间
func MonthBefore(tm time.Time, n int) time.Time {
	// 本月
	//m := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, tm.Location())
	month := tm.Year()*12 + int(tm.Month()) - 1 - n
	return time.Date(month/12, time.Month(month%12+1), 1, 0, 0, 0, 0, tm.Location())
}

// YearBefore n年前
func YearBefore(tm time.Time, n int, yearStart bool) time.Time {
	if yearStart {
		return time.Date(tm.Year()-n, 1, 1, 0, 0, 0, 0, tm.Location())
	}
	return time.Date(tm.Year()-n, tm.Month(), tm.Day(),
		tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond(), tm.Location())
}

// Tomorrow 比当前时间晚1个自然天的自然天的开始时间
func Tomorrow(tm time.Time) (ret time.Time) {
	return After(tm, 1)
}

// Yesterday 比当前时间早1个自然天的自然天的开始时间
func Yesterday(tm time.Time) (ret time.Time) {
	return Before(tm, 1)
}

// today 今天开始的时间，当前时区
func today(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

// Today 返回本自然天的开始时间
func Today() (ret time.Time) {
	return today(time.Now())
}

// Days 自从1970-1-1到tm的天数
func Days(tm time.Time) int {
	_, offset := tm.Zone()
	return int((tm.Unix() + int64(offset)) / int64(86400))
}

// Between 两个时间之间的自然天差值
func Between(t1, t2 time.Time) (days int) {
	d1 := Days(t1)
	d2 := Days(t2)
	if d1 > d2 {
		return d1 - d2
	}
	return d2 - d1
}

// After N天以后的自然天的起始时间
func After(tm time.Time, n int) (ret time.Time) {
	t := today(tm)
	return time.Unix(t.Unix()+int64(86400*n), 0)
}

// Before N天之前的自然天的起始时间
func Before(tm time.Time, n int) (ret time.Time) {
	t := today(tm)
	return time.Unix(t.Unix()-int64(86400*n), 0)
}
