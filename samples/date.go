package samples

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/date"
)

func dateTime() {
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

func dateISO() {
	println("ISO:", date.ISO())
}

func dateShort() {
	println("Short:", date.Short())
	println("ShortNormal:", date.ShortNormal())
	println("ShortReverse:", date.ShortReverse())
	println("ShortDash:", date.ShortDash())
	println("ShortDot:", date.ShortDot())
	println("ShortUnderscore:", date.ShortUnderscore())
	println("ShortSpace:", date.ShortSpace())
	println("ShortMonth:", date.ShortMonth())
}

func dateLong() {
	println("Long:", date.Long())
}

func dateUTC() {
	println("UTC:", date.UTC())
}

func dateToSecond() {
	println("ToSecond:", date.ToSecond(1))
}

func dateToMinute() {
	println("ToMinute:", date.ToMinute(1))
}

func dateToHour() {
	println("ToHour:", date.ToHour(1))
}

func dateSleepSecond() {
	date.SleepSecond(1)
}

func dateSleepMinute() {
	date.SleepMinute(1)
}

func dateSleepHour() {
	date.SleepHour(1)
}

func dateClock() {
	date.TerminalClock()
}
