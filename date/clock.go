package date

import (
	"fmt"
	"strconv"
	"time"

	"github.com/thuongtruong1009/gouse/console"
)

func formatTime(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secondStr := strconv.Itoa(second)

	if hour < 10 {
		hourStr = fmt.Sprintf("0%s", hourStr)
	}
	if minute < 10 {
		minuteStr = fmt.Sprintf("0%s", minuteStr)
	}
	if second < 10 {
		secondStr = fmt.Sprintf("0%s", secondStr)
	}

	return fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}

func TerminalClock() {
	msgTime := make(chan time.Time)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			msgTime <- time.Now()
		}
	}()

	for t := range msgTime {
		console.Clear()
		fmt.Println(formatTime(t))
	}
}
