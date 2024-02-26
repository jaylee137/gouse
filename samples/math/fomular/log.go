package fomular

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathLog() {
	println("Logarithm of integer number: ", math.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", math.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", math.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", math.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", math.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", math.Log10F(20.0)))
}
