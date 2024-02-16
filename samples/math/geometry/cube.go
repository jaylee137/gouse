package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathCube() {
	println("Area of cube (integer): ", math.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", math.AreaCubeF(10.0)))
	println("Perimeter of cube (integer): ", math.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", math.PeriCubeF(10.0)))
	println("Volume of cube (integer): ", math.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", math.VolCubeF(10.0)))
}
