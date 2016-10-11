package days

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	fmt.Println(time.Now(), "Timer should stop after 5 second.")
	tmr := StartTimer(1000, 2000, func(args ...interface{}) error {
		fmt.Println("timer: ", time.Now())
		return nil
	})

	go func() {
		time.Sleep(5 * time.Second)
		StopTimer(tmr)
	}()

	time.Sleep(10 * time.Second)
}
