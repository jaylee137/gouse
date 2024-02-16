package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathCircle() {
	println("Area of circle (integer): ", fmt.Sprintf("%f", math.AreaCircle(10)))
	println("Area of circle (float): ", fmt.Sprintf("%f", math.AreaCircleF(10.0)))
	println("Perimeter of circle (integer): ", fmt.Sprintf("%f", math.PeriCircle(10)))
	println("Perimeter of circle (float): ", fmt.Sprintf("%f", math.PeriCircleF(10.0)))
}
