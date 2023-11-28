package examples

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
	println(date.ISO())
}

func dateShort() {
	println(date.Short())
}

func dateShortInverse() {
	println(date.ShortInverse())
}

func dateLong() {
	println(date.Long())
}

// func dateCustom() {
// 	println(date.Custom("2006-01-02", "02/01/2006"))
// }

func DateExample() {
	date.Ok()

	dateTime()
	dateISO()
	dateShort()
	dateShortInverse()
	dateLong()
	// dateCustom()
}