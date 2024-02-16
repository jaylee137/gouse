package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathSphere() {
	println("Area of sphere (integer): ", fmt.Sprintf("%f", math.AreaSphere(10)))
	println("Area of sphere (float): ", fmt.Sprintf("%f", math.AreaSphereF(10.0)))
	println("Volume of sphere (integer): ", fmt.Sprintf("%f", math.VolSphere(10)))
	println("Volume of sphere (float): ", fmt.Sprintf("%f", math.VolSphereF(10.0)))
}
