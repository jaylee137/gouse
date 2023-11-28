package function

import "github.com/thuongtruong1009/gouse/date"

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
