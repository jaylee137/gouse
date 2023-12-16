package date

import (
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	if Second() != time.Now().Second() {
		t.Error("Second() != time.Now().Second()")
	}
}

func TestMinute(t *testing.T) {
	if Minute() != time.Now().Minute() {
		t.Error("Minute() != time.Now().Minute()")
	}
}

func TestHour(t *testing.T) {
	if Hour() != time.Now().Hour() {
		t.Error("Hour() != time.Now().Hour()")
	}
}

func TestDay(t *testing.T) {
	if Day() != time.Now().Day() {
		t.Error("Day() != time.Now().Day()")
	}
}

func TestMonth(t *testing.T) {
	if Month() != int(time.Now().Month()) {
		t.Error("Month() != int(time.Now().Month())")
	}
}

func TestYear(t *testing.T) {
	if Year() != time.Now().Year() {
		t.Error("Year() != time.Now().Year()")
	}
}

func TestWeekday(t *testing.T) {
	if Weekday() != int(time.Now().Weekday()) {
		t.Error("Weekday() != int(time.Now().Weekday())")
	}
}

func TestUnix(t *testing.T) {
	if Unix() != time.Now().Unix() {
		t.Error("Unix() != time.Now().Unix()")
	}
}

func TestUnixMilli(t *testing.T) {
	if UnixMilli() != time.Now().UnixNano()/int64(time.Millisecond) {
		t.Error("UnixMilli() != time.Now().UnixNano()/int64(time.Millisecond)")
	}
}

func TestUnixMicro(t *testing.T) {
	result := UnixMicro()
	currentTime := time.Now().UnixNano() / int64(time.Microsecond)

	acceptableRange := int64(1000) // 1 millisecond in microseconds

	if result < currentTime-acceptableRange || result > currentTime+acceptableRange {
		t.Errorf("UnixMicro() result %d is not within an acceptable range of current time %d", result, currentTime)
	}
}

func TestUnixNano(t *testing.T) {
	result := UnixNano()
	currentTime := time.Now().UnixNano()

	acceptableRange := int64(1000000000) // 1 second in nanoseconds

	if result < currentTime-acceptableRange || result > currentTime+acceptableRange {
		t.Errorf("UnixNano() result %d is not within an acceptable range of current time %d", result, currentTime)
	}
}

func TestUnixMilliToTime(t *testing.T) {
	if UnixMilliToTime(0) != time.Unix(0, 0) {
		t.Error("UnixMilliToTime(0) != time.Unix(0, 0)")
	}
}

func TestUnixMicroToTime(t *testing.T) {
	if UnixMicroToTime(0) != time.Unix(0, 0) {
		t.Error("UnixMicroToTime(0) != time.Unix(0, 0)")
	}
}

func TestUnixNanoToTime(t *testing.T) {
	if UnixNanoToTime(0) != time.Unix(0, 0) {
		t.Error("UnixNanoToTime(0) != time.Unix(0, 0)")
	}
}
