package operator

import "github.com/thuongtruong1009/gouse/math"

func SampleMathSum() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", math.Sum(num1, num2, num3, num4))
}
