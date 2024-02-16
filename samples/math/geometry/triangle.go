package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathTriangle() {
	println("Area of triangle (integer): ", math.AreaTriangle(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", math.AreaTriangleF(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", math.PeriTriangle(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", math.PeriTriangleF(10.0, 20.0, 30.0)))
}
