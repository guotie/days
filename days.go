package days

import (
	"time"
)

// 比当前时间晚1个自然天的自然天的开始时间
func Tomorrow(tm time.Time) (ret time.Time) {
	return DaysAfter(tm, 1)
}

// 比当前时间早1个自然天的自然天的开始时间
func Yesterday(tm time.Time) (ret time.Time) {
	return DaysBefore(tm, 1)
}

func today(tm time.Time) (ret time.Time) {
	_, offset := tm.Zone()
	seconds := tm.Unix()
	return time.Unix(seconds-(seconds+int64(offset))%86400, 0)
}

// 返回本自然天的开始时间
func Today() (ret time.Time) {
	return today(time.Now())
}

func Days(tm time.Time) int {
	_, offset := tm.Zone()
	return int((tm.Unix() + int64(offset)) / int64(86400))
}

// 两个时间之间的自然天差值
func DaysBetween(t1, t2 time.Time) (days int) {
	d1 := Days(t1)
	d2 := Days(t2)
	if d1 > d2 {
		return d1 - d2
	}
	return d2 - d1
}

// N天以后的自然天的起始时间
func DaysAfter(tm time.Time, n int) (ret time.Time) {
	t := today(tm)
	return t.Add(time.Duration(n) * time.Duration(86400))
}

// N天之前的自然天的起始时间
func DaysBefore(tm time.Time, n int) (ret time.Time) {
	t := today(tm)
	return t.Add(time.Duration(-n) * time.Duration(86400))
}
