package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathSquare() {
	println("Area of square (integer): ", math.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", math.AreaSquareF(10.0)))
	println("Perimeter of square (integer): ", math.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", math.PeriSquareF(10.0)))
}
