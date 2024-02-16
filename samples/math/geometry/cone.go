package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathCone() {
	println("Area of cone (integer): ", fmt.Sprintf("%f", math.AreaCone(10, 20)))
	println("Area of cone (float): ", fmt.Sprintf("%f", math.AreaConeF(10.0, 20.0)))
	println("Volume of cone (integer): ", fmt.Sprintf("%f", math.VolCone(10, 20)))
	println("Volume of cone (float): ", fmt.Sprintf("%f", math.VolConeF(10.0, 20.0)))
}
