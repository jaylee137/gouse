package operator

import "github.com/thuongtruong1009/gouse/math"

func mathMin() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Min number: ", math.Min(num1, num2, num3, num4))
}
