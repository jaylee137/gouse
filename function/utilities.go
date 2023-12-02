package function

import (
	"time"

	"github.com/thuongtruong1009/gouse/date"
)

func Retry(fn func() error, attempts int, sleep int) (err error) {
	retry := func() {
		if err = fn(); err != nil {
			return
		}
	}

	for i := attempts; i > 0; i-- {
		if i > 1 {
			retry()
			date.SleepSecond(sleep)
		} else if i == 1 {
			retry()
		} else {
			return nil
		}
	}
	return
}

func Times(fn func(), attempts int) {
	for i := 0; i < attempts; i++ {
		fn()
	}
}

func Interval(fn func(), timeout int) {
	for {
		fn()
		date.SleepSecond(timeout)
	}
}

func RunTime(startTime time.Time, task func()) time.Duration {
	task()
	elapsedTime := float64(time.Since(startTime).Seconds() * 1000)
	return time.Duration(elapsedTime)
}
