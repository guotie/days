package days

import (
	"time"
)

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
