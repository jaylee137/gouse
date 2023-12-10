package date

import (
	"testing"
	"time"
)

func TestISO(t *testing.T) {
	now := time.Now()
	if ISO() != now.Format("2006-01-02T15:04:05.999Z") {
		t.Error("ISO() should return today's date in ISO format")
	}
}

func TestShort(t *testing.T) {
	if Short() != time.Now().Format("02/01/2006") {
		t.Error("Short() should return today's date in Short format")
	}
}

func TestShortNormal(t *testing.T) {
	if ShortNormal() != time.Now().Format("01/02/2006") {
		t.Error("ShortNormal() should return today's date in ShortNormal format")
	}
}

func TestShortReverse(t *testing.T) {
	if ShortReverse() != time.Now().Format("2006/01/02") {
		t.Error("ShortReverse() should return today's date in ShortReverse format")
	}
}

func TestShortDash(t *testing.T) {
	if ShortDash() != time.Now().Format("2006-01-02") {
		t.Error("ShortDash() should return today's date in ShortDash format")
	}
}

func TestShortDot(t *testing.T) {
	if ShortDot() != time.Now().Format("2006.01.02") {
		t.Error("ShortDot() should return today's date in ShortDot format")
	}
}

func TestShortUnderscore(t *testing.T) {
	if ShortUnderscore() != time.Now().Format("2006_01_02") {
		t.Error("ShortUnderscore() should return today's date in ShortUnderscore format")
	}
}

func TestShortSpace(t *testing.T) {
	if ShortSpace() != time.Now().Format("2006 01 02") {
		t.Error("ShortSpace() should return today's date in ShortSpace format")
	}
}

func TestShortMonth(t *testing.T) {
	if ShortMonth() != time.Now().Format("01/2006") {
		t.Error("ShortMonth() should return today's date in ShortMonth format")
	}
}

func TestLong(t *testing.T) {
	if Long() != time.Now().Format("Monday, January 02, 2006") {
		t.Error("Long() should return today's date in Long format")
	}
}

func TestUTC(t *testing.T) {
	if UTC() != time.Now().Format("Jan 2, 2006 at 3:04pm (MST)") {
		t.Error("UTC() should return today's date in UTC format")
	}
}
