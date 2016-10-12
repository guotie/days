package days

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	fmt.Println(time.Now(), "Timer should stop after 5 minute.")
	tmr := StartTimer(1000, 60*1000, func(args ...interface{}) error {
		fmt.Println("timer: ", time.Now())
		return nil
	})

	go func() {
		time.Sleep(5 * time.Minute)
		StopTimer(tmr)
	}()

	time.Sleep(10 * time.Minute)
}
