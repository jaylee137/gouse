package date

import (
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	if Second() != time.Now().Second() {
		t.Fail()
	}
}

func TestMinute(t *testing.T) {
	if Minute() != time.Now().Minute() {
		t.Fail()
	}
}

func TestHour(t *testing.T) {
	if Hour() != time.Now().Hour() {
		t.Fail()
	}
}

func TestDay(t *testing.T) {
	if Day() != time.Now().Day() {
		t.Fail()
	}
}

func TestMonth(t *testing.T) {
	if Month() != int(time.Now().Month()) {
		t.Fail()
	}
}

func TestYear(t *testing.T) {
	if Year() != time.Now().Year() {
		t.Fail()
	}
}

func TestWeekday(t *testing.T) {
	if Weekday() != int(time.Now().Weekday()) {
		t.Fail()
	}
}

func TestUnix(t *testing.T) {
	if Unix() != time.Now().Unix() {
		t.Fail()
	}
}

func TestUnixMilli(t *testing.T) {
	if UnixMilli() != time.Now().UnixNano()/int64(time.Millisecond) {
		t.Fail()
	}
}

func TestUnixMicro(t *testing.T) {
	if UnixMicro() != time.Now().UnixNano()/int64(time.Microsecond) {
		t.Fail()
	}
}

func TestUnixNano(t *testing.T) {
	if UnixNano() != time.Now().UnixNano() {
		t.Fail()
	}
}

func TestUnixMilliToTime(t *testing.T) {
	if UnixMilliToTime(0) != time.Unix(0, 0) {
		t.Fail()
	}
}

func TestUnixMicroToTime(t *testing.T) {
	if UnixMicroToTime(0) != time.Unix(0, 0) {
		t.Fail()
	}
}

func TestUnixNanoToTime(t *testing.T) {
	if UnixNanoToTime(0) != time.Unix(0, 0) {
		t.Fail()
	}
}
