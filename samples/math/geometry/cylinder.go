package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathCylinder() {
	println("Area of cylinder (integer): ", fmt.Sprintf("%f", math.AreaCylinder(10, 20)))
	println("Area of cylinder (float): ", fmt.Sprintf("%f", math.AreaCylinderF(10.0, 20.0)))
	println("Volume of cylinder (integer): ", fmt.Sprintf("%f", math.VolCylinder(10, 20)))
	println("Volume of cylinder (float): ", fmt.Sprintf("%f", math.VolCylinderF(10.0, 20.0)))
}
