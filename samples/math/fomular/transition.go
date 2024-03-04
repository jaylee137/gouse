package fomular

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathTransition() {
	var distance, speed, time float64 = 100, 10, 10
	println("Speed: ", fmt.Sprintf("%f", math.Speed(distance, time)))
	println("Distance: ", fmt.Sprintf("%f", math.Distance(speed, time)))
	println("Time: ", fmt.Sprintf("%f", math.Time(distance, speed)))
}
