package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathPolygon() {
	println("Area of pentagol by number of sides (integer): ", fmt.Sprintf("%f", math.AreaPolygon(10, 6)))
}
