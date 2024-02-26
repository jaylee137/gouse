package fomular

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathTrigonometry() {
	println("Sine of integer number: ", math.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", math.SinF(90.0)))
	println("Cosine of integer number: ", math.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", math.CosF(90.0)))
	println("Tangent of integer number: ", math.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", math.TanF(90.0)))
}
