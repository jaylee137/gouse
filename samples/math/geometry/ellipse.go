package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathEllipse() {
	println("Area of ellipse (integer): ", fmt.Sprintf("%f", math.AreaEllipse(10, 20)))
	println("Area of ellipse (float): ", fmt.Sprintf("%f", math.AreaEllipseF(10.0, 20.0)))
}
