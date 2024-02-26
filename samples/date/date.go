package date

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/date"
)

func SampleDateTime() {
	println("Second:", date.Second())
	println("Minute:", date.Minute())
	println("Hour:", date.Hour())
	println("Day:", date.Day())
	println("Month:", date.Month())
	println("Year:", date.Year())
	println("Weekday:", date.Weekday())
	println("Unix:", date.Unix())
	println("UnixNano:", date.UnixNano())
	println("UnixMilli:", date.UnixMilli())
	println("UnixMicro:", date.UnixMicro())
	fmt.Println("UnixMilliToTime:", date.UnixMilliToTime(1000000000))
	fmt.Println("UnixMicroToTime:", date.UnixMicroToTime(1000000000))
	fmt.Println("UnixNanoToTime:", date.UnixNanoToTime(1000000000))
}

func SampleDateISO() {
	println("ISO:", date.ISO())
}

func SampleDateShort() {
	println("Short:", date.Short())
	println("ShortNormal:", date.ShortNormal())
	println("ShortReverse:", date.ShortReverse())
	println("ShortDash:", date.ShortDash())
	println("ShortDot:", date.ShortDot())
	println("ShortUnderscore:", date.ShortUnderscore())
	println("ShortSpace:", date.ShortSpace())
	println("ShortMonth:", date.ShortMonth())
}

func SampleDateLong() {
	println("Long:", date.Long())
}

func SampleDateUTC() {
	println("UTC:", date.UTC())
}

func SampleDateToSecond() {
	println("ToSecond:", date.ToSecond(1))
}

func SampleDateToMinute() {
	println("ToMinute:", date.ToMinute(1))
}

func SampleDateToHour() {
	println("ToHour:", date.ToHour(1))
}

func SampleDateSleepSecond() {
	date.SleepSecond(1)
}

func SampleDateSleepMinute() {
	date.SleepMinute(1)
}

func SampleDateSleepHour() {
	date.SleepHour(1)
}

func SampleDateClock() {
	date.TerminalClock()
}
