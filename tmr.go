package days

import (
	"time"
)

// 与uv_timer_start类似
// uv_timer_start(tmr, first run time, loop time, fn)
func StartTimer(n, m int64, fn func(args ...interface{}) error, args ...interface{}) *time.Timer {
	tmr := time.NewTimer(time.Duration(n) * time.Millisecond)

	go func() {
		<-tmr.C
		fn(args)

		if m != 0 {
			for {
				tmr.Reset(time.Duration(m) * time.Millisecond)
				<-tmr.C
				fn(args)
			}
		}
	}()

	return tmr
}
