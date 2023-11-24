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