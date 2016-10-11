package days

import (
	"time"
)

// StartTimer 与uv_timer_start类似
// uv_timer_start(tmr, first run time, loop time, fn)
// n: 毫秒
// m: 毫秒
func StartTimer(n, m int64, fn func(args ...interface{}) error, args ...interface{}) *time.Timer {
	tmr := time.NewTimer(time.Duration(n) * time.Millisecond)

	go func() {
		<-tmr.C
		fn(args)

		if m != 0 {
			for {
				// 修复时间偏差
				// 仅限小时定时和天定时
				seconds := m
				now := time.Now()
				today := Today()
				switch m {
				case 3600:
					seconds -= now.Unix() % 3600
				case 86400:
					seconds -= int64(now.Sub(today).Seconds())
				}

				tmr.Reset(time.Duration(m) * time.Millisecond)
				<-tmr.C
				fn(args)
			}
		}
	}()

	return tmr
}
