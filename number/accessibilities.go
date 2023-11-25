package number

import (
	"time"
	"math/rand"
)

func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)

	return min + rand.Intn(max-min)
}

func Clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}

func InRange(value int, min int, max int) bool {
	return value >= min && value <= max
}