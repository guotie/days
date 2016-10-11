package days

import (
	"time"
)

// Timer 定时器
type Timer struct {
	*time.Timer
	exitCh chan bool
}

// StartTimer 启动定时器
// 与uv_timer_start类似
// uv_timer_start(tmr, first run time, loop time, fn)
// n: 毫秒
// m: 毫秒
func StartTimer(n, m int64, fn func(args ...interface{}) error, args ...interface{}) *Timer {
	tmr := time.NewTimer(time.Duration(n) * time.Millisecond)
	exitCh := make(chan bool, 0)

	doTimer := func() bool {
		select {
		case <-tmr.C:
		case <-exitCh:
			// 定时器退出
			tmr.Stop()
			return true
		}
		// 执行定时器函数
		fn(args)
		return false
	}

	go func() {
		if doTimer() == true {
			return
		}

		if m == 0 {
			return
		}

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

			// 重置定时器
			tmr.Reset(time.Duration(m) * time.Millisecond)
			if doTimer() == true {
				return
			}
		}
	}()

	return &Timer{tmr, exitCh}
}

// StopTimer 停止定时器
func StopTimer(tmr *Timer) {
	tmr.exitCh <- true
}
