package date

import "time"

func ToSecond(second int) time.Duration {
	return time.Duration(second) * time.Second
}

func ToMinute(minute int) time.Duration {
	return time.Duration(minute) * time.Minute
}

func ToHour(hour int) time.Duration {
	return time.Duration(hour) * time.Hour
}

func SleepSecond(second int) {
	time.Sleep(ToSecond(second))
}

func SleepMinute(minute int) {
	time.Sleep(ToMinute(minute))
}

func SleepHour(hour int) {
	time.Sleep(ToHour(hour))
}