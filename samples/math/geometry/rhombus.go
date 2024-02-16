package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathRhombus() {
	println("Area of rhombus (integer): ", math.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", math.AreaRhombusF(10.0, 20.0)))
}
