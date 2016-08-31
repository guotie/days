package days

import (
	"testing"
	"time"
)

var tms = []time.Time{
	time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 0, 0, 0, 1, time.Local),
	time.Date(2015, time.Month(9), 22, 7, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 7, 0, 1, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 8, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 8, 0, 1, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 15, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 15, 59, 1, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 16, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 16, 59, 1, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 23, 0, 0, 0, time.Local),
	time.Date(2015, time.Month(9), 22, 23, 59, 1, 0, time.Local),
}

var daysdelta = []struct {
	t1    time.Time
	t2    time.Time
	delta int
}{
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 7, 0, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 8, 0, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 15, 0, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 16, 0, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 22, 23, 59, 0, 1, time.Local),
		0,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 0, 0, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 7, 0, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 8, 0, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 15, 0, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 16, 0, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 23, 23, 59, 0, 1, time.Local),
		1,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 0, 0, 0, 1, time.Local),
		2,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 7, 0, 0, 1, time.Local),
		2,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 8, 0, 0, 1, time.Local),
		2,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 15, 0, 0, 1, time.Local),
		2,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 16, 0, 0, 1, time.Local),
		2,
	},
	{
		time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local),
		time.Date(2015, time.Month(9), 24, 23, 59, 0, 1, time.Local),
		2,
	},
}

func TestToday(t *testing.T) {
	t0 := time.Date(2015, time.Month(9), 22, 0, 0, 0, 0, time.Local)
	for _, tm := range tms {
		if today(tm) != t0 {
			t.Logf("today of %v is %v, but should be %v\n", tm, today(tm), t0)
			t.Fail()
		}
	}

}

func TestDays(t *testing.T) {
	d0 := Days(tms[0])
	for _, tm := range tms {
		if Days(tm) != d0 {
			t.Fail()
		}
	}
}

func TestDaysBetween(t *testing.T) {
	for _, dd := range daysdelta {
		if DaysBetween(dd.t1, dd.t2) != dd.delta {
			t.Fail()
		}
	}
}

func TestDaysAfter(t *testing.T) {
	ret := time.Date(2015, time.Month(9), 23, 0, 0, 0, 0, time.Local)
	for idx, tm := range tms {
		if DaysAfter(tm, 1) != ret {
			t.Logf("Before time %v 1 days should be %v, but %v, index %d\n", tm, ret, DaysAfter(tm, 1), idx)
			t.Fail()
		}
	}
}
func TestDaysBefore(t *testing.T) {
	ret := time.Date(2015, time.Month(9), 21, 0, 0, 0, 0, time.Local)
	for idx, tm := range tms {
		if DaysBefore(tm, 1) != ret {
			t.Logf("Before time %v 1 days should be %v, but %v, index %d\n", tm, ret, DaysBefore(tm, 1), idx)
			t.Fail()
		}
	}
}
