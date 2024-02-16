package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathRect() {
	println("Area of rectangle: ", math.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", math.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", math.PeriRectF(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%f", math.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", math.DiagRectF(10.0, 20.0)))
	println("Volume of rectangular (integer): ", math.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", math.VolRectF(10.0, 20.0, 30.0)))
}
