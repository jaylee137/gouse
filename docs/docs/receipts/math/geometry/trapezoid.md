# Trapezoid

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/math"
)
```
## Functions


### SampleMathTrapezoid

```go
func SampleMathTrapezoid() {
	println("Area of trapezoid (integer): ", fmt.Sprintf("%f", math.AreaTrapezoid(10, 20, 30)))
	println("Area of trapezoid (float): ", fmt.Sprintf("%f", math.AreaTrapezoidF(10.0, 20.0, 30.0)))
}
```
