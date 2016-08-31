package days

import (
	"fmt"
	"testing"
	"time"
)

func TestStartTimer(t *testing.T) {
	fmt.Println(time.Now())
	StartTimer(1000, 2000, func(args ...interface{}) error {
		fmt.Println(time.Now())
		return nil
	})

	time.Sleep(10 * time.Second)
}
