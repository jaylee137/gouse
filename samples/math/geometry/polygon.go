package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func mathPolygon() {
	println("Area of pentagol by number of sides (integer): ", fmt.Sprintf("%f", math.AreaPolygon(10, 6)))
}
