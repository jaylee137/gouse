package date

import "time"

func Second() int {
	return time.Now().Second()
}

func Minute() int {
	return time.Now().Minute()
}

func Hour() int {
	return time.Now().Hour()
}

func Day() int {
	return time.Now().Day()
}

func Month() int {
	return int(time.Now().Month())
}

func Year() int {
	return time.Now().Year()
}

func Weekday() int {
	return int(time.Now().Weekday())
}

func Unix() int64 {
	return time.Now().Unix()
}

func UnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func UnixMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}

func UnixMicroToTime(micro int64) time.Time {
	return time.Unix(0, micro*int64(time.Microsecond))
}

func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(0, nano)
}
