# Operator


### IsPrime
```go
import ()
```

```go
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
```

### IsEven
```go
import ()
```

```go
func IsEven(num int) bool {
	return num%2 == 0
}
```

### IsOdd
```go
import ()
```

```go
func IsOdd(num int) bool {
	return num%2 != 0
}
```

### IsPerfectSquare
```go
import ()
```

```go
func IsPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
```

### Log
```go
import (
	"math"
)
```

```go
func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}
```

### LogF
```go
import (
	"math"
)
```

```go
func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}
```

### Log2
```go
import (
	"math"
)
```

```go
func Log2(number int) int {
	return Log(number, 2)
}
```

### Log2F
```go
import (
	"math"
)
```

```go
func Log2F(number float64) float64 {
	return LogF(number, 2)
}
```

### Log10
```go
import (
	"math"
)
```

```go
func Log10(number int) int {
	return Log(number, 10)
}
```

### Log10F
```go
import (
	"math"
)
```

```go
func Log10F(number float64) float64 {
	return LogF(number, 10)
}
```

### Pytago
```go
import (
	"math"
)
```

```go
func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}
```

### PytagoF
```go
import (
	"math"
)
```

```go
func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}
```

### Sin
```go
import (
	"math"
)
```

```go
func Sin(angle int) float64 {
	return math.Sin(float64(angle))
}
```

### SinF
```go
import (
	"math"
)
```

```go
func SinF(angle float64) float64 {
	return math.Sin(angle)
}
```

### Cos
```go
import (
	"math"
)
```

```go
func Cos(angle int) float64 {
	return math.Cos(float64(angle))
}
```

### CosF
```go
import (
	"math"
)
```

```go
func CosF(angle float64) float64 {
	return math.Cos(angle)
}
```

### Tan
```go
import (
	"math"
)
```

```go
func Tan(angle int) float64 {
	return math.Tan(float64(angle))
}
```

### TanF
```go
import (
	"math"
)
```

```go
func TanF(angle float64) float64 {
	return math.Tan(angle)
}
```

### Speed
```go
import (
	"math"
)
```

```go
func Speed(distance, time float64) float64 {
	return DivideF(distance, time)
}
```

### Distance
```go
import (
	"math"
)
```

```go
func Distance(speed, time float64) float64 {
	return MultiF(speed, time)
}
```

### Time
```go
import (
	"math"
)
```

```go
func Time(distance, speed float64) float64 {
	return DivideF(distance, speed)
}
```

### AreaRect
```go
import (
	"math"
)
```

```go
func AreaRect(length, width int) int {
	return Multi(length, width)
}
```

### AreaRectF
```go
import (
	"math"
)
```

```go
func AreaRectF(length, width float64) float64 {
	return MultiF(length, width)
}
```

### PeriRect
```go
import (
	"math"
)
```

```go
func PeriRect(length, width int) int {
	return Multi(2, Add(length, width))
}
```

### PeriRectF
```go
import (
	"math"
)
```

```go
func PeriRectF(length, width float64) float64 {
	return MultiF(2, AddF(length, width))
}
```

### DiagRect
```go
import (
	"math"
)
```

```go
func DiagRect(length, width int) float64 {
	return Pytago(length, width)
}
```

### DiagRectF
```go
import (
	"math"
)
```

```go
func DiagRectF(length, width float64) float64 {
	return PytagoF(length, width)
}
```

### VolRect
```go
import (
	"math"
)
```

```go
func VolRect(length, width, height int) int {
	return Multi(length, width, height)
}
```

### VolRectF
```go
import (
	"math"
)
```

```go
func VolRectF(length, width, height float64) float64 {
	return MultiF(length, width, height)
}
```

### AreaCircle
```go
import (
	"math"
)
```

```go
func AreaCircle(radius int) float64 {
	return math.Pi * float64(Pow2(radius))
}
```

### AreaCircleF
```go
import (
	"math"
)
```

```go
func AreaCircleF(radius float64) float64 {
	return math.Pi * PowF(radius, 2)
}
```

### PeriCircle
```go
import (
	"math"
)
```

```go
func PeriCircle(radius int) float64 {
	return 2 * math.Pi * float64(radius)
}
```

### PeriCircleF
```go
import (
	"math"
)
```

```go
func PeriCircleF(radius float64) float64 {
	return 2 * math.Pi * radius
}
```

### AreaTriangle
```go
import (
	"math"
)
```

```go
func AreaTriangle(base, height int) int {
	return Divide(Multi(base, height), 2)
}
```

### AreaTriangleF
```go
import (
	"math"
)
```

```go
func AreaTriangleF(base, height float64) float64 {
	return DivideF(MultiF(base, height), 2)
}
```

### PeriTriangle
```go
import (
	"math"
)
```

```go
func PeriTriangle(side1, side2, side3 int) int {
	return Sum(side1, side2, side3)
}
```

### PeriTriangleF
```go
import (
	"math"
)
```

```go
func PeriTriangleF(side1, side2, side3 float64) float64 {
	return SumF(side1, side2, side3)
}
```

### AreaSquare
```go
import (
	"math"
)
```

```go
func AreaSquare(side int) int {
	return Pow2(side)
}
```

### AreaSquareF
```go
import (
	"math"
)
```

```go
func AreaSquareF(side float64) float64 {
	return PowF(side, 2)
}
```

### PeriSquare
```go
import (
	"math"
)
```

```go
func PeriSquare(side int) int {
	return Multi(side, 4)
}
```

### PeriSquareF
```go
import (
	"math"
)
```

```go
func PeriSquareF(side float64) float64 {
	return MultiF(side, 4)
}
```

### AreaCube
```go
import (
	"math"
)
```

```go
func AreaCube(side int) int {
	return Multi(Pow2(side), 6)
}
```

### AreaCubeF
```go
import (
	"math"
)
```

```go
func AreaCubeF(side float64) float64 {
	return MultiF(PowF(side, 2), 6)
}
```

### PeriCube
```go
import (
	"math"
)
```

```go
func PeriCube(side int) int {
	return Multi(side, 12)
}
```

### PeriCubeF
```go
import (
	"math"
)
```

```go
func PeriCubeF(side float64) float64 {
	return MultiF(side, 12)
}
```

### VolCube
```go
import (
	"math"
)
```

```go
func VolCube(side int) int {
	return Pow3(side)
}
```

### VolCubeF
```go
import (
	"math"
)
```

```go
func VolCubeF(side float64) float64 {
	return PowF(side, 3)
}
```

### AreaSphere
```go
import (
	"math"
)
```

```go
func AreaSphere(radius int) float64 {
	return 4 * math.Pi * float64(Pow2(radius))
}
```

### AreaSphereF
```go
import (
	"math"
)
```

```go
func AreaSphereF(radius float64) float64 {
	return 4 * math.Pi * PowF(radius, 2)
}
```

### VolSphere
```go
import (
	"math"
)
```

```go
func VolSphere(radius int) float64 {
	return (4 / 3) * math.Pi * float64(Pow2(radius))
}
```

### VolSphereF
```go
import (
	"math"
)
```

```go
func VolSphereF(radius float64) float64 {
	return (4 / 3) * math.Pi * PowF(radius, 2)
}
```

### AreaCylinder
```go
import (
	"math"
)
```

```go
func AreaCylinder(radius, height int) float64 {
	return (2 * math.Pi * float64(radius)) * float64(height)
}
```

### AreaCylinderF
```go
import (
	"math"
)
```

```go
func AreaCylinderF(radius, height float64) float64 {
	return (2 * math.Pi * radius) * height
}
```

### VolCylinder
```go
import (
	"math"
)
```

```go
func VolCylinder(radius, height int) float64 {
	return math.Pi * float64(Pow2(radius)) * float64(height)
}
```

### VolCylinderF
```go
import (
	"math"
)
```

```go
func VolCylinderF(radius, height float64) float64 {
	return math.Pi * PowF(radius, 2) * height
}
```

### AreaCone
```go
import (
	"math"
)
```

```go
func AreaCone(radius, height int) float64 {
	return math.Pi * float64(radius) * (float64(radius) + math.Sqrt(float64(Pow2(height)+Pow2(radius))))
}
```

### AreaConeF
```go
import (
	"math"
)
```

```go
func AreaConeF(radius, height float64) float64 {
	return math.Pi * radius * (radius + math.Sqrt(PowF(height, 2)+PowF(radius, 2)))
}
```

### VolCone
```go
import (
	"math"
)
```

```go
func VolCone(radius, height int) float64 {
	return DivideF(1, 3) * math.Pi * float64(Pow2(radius)) * float64(height)
}
```

### VolConeF
```go
import (
	"math"
)
```

```go
func VolConeF(radius, height float64) float64 {
	return DivideF(1, 3) * math.Pi * PowF(radius, 2) * height
}
```

### AreaTrapezoid
```go
import (
	"math"
)
```

```go
func AreaTrapezoid(base1, base2, height int) float64 {
	return 0.5 * float64(Add(base1, base2)) * float64(height)
}
```

### AreaTrapezoidF
```go
import (
	"math"
)
```

```go
func AreaTrapezoidF(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}
```

### AreaParallelogram
```go
import (
	"math"
)
```

```go
func AreaParallelogram(base, height int) int {
	return Multi(base, height)
}
```

### AreaParallelogramF
```go
import (
	"math"
)
```

```go
func AreaParallelogramF(base, height float64) float64 {
	return MultiF(base, height)
}
```

### AreaRhombus
```go
import (
	"math"
)
```

```go
func AreaRhombus(diag1, diag2 int) int {
	return Divide(Multi(diag1, diag2), 2)
}
```

### AreaRhombusF
```go
import (
	"math"
)
```

```go
func AreaRhombusF(diag1, diag2 float64) float64 {
	return DivideF(MultiF(diag1, diag2), 2)
}
```

### AreaEllipse
```go
import (
	"math"
)
```

```go
func AreaEllipse(major, minor int) float64 {
	return math.Pi * float64(major) * float64(minor)
}
```

### AreaEllipseF
```go
import (
	"math"
)
```

```go
func AreaEllipseF(major, minor float64) float64 {
	return math.Pi * major * minor
}
```

### AreaPolygon
```go
import (
	"math"
)
```

```go
func AreaPolygon(lenSide float64, numSide int) float64 {
	return DivideF(MultiF(0.25, float64(numSide), Pow2F(lenSide)), AbsF(TanF(DivideF(180, float64(numSide)))))
}
```

### Abs
```go
import (
	"math"
)
```

```go
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
```

### AbsF
```go
import (
	"math"
)
```

```go
func AbsF(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}
```

### Floor
```go
import (
	"math"
)
```

```go
func Floor(num float64) int {
	return int(num)
}
```

### Ceil
```go
import (
	"math"
)
```

```go
func Ceil(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	if num > 0 {
		return int(num) + 1
	}
	return int(num) - 1
}
```

### Round
```go
import (
	"math"
)
```

```go
func Round(num float64) int {
	if num < 0 {
		return int(num)
	}

	if num-float64(int(num)) < 0.5 {
		return int(num)
	}

	return int(num) + 1
}
```

### Max
```go
import (
	"math"
)
```

```go
func Max(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
```

### MaxF
```go
import (
	"math"
)
```

```go
func MaxF(nums ...float64) float64 {
	var max float64
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
```

### Min
```go
import (
	"math"
)
```

```go
func Min(nums ...int) int {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}
```

### MinF
```go
import (
	"math"
)
```

```go
func MinF(nums ...float64) float64 {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}
```

### Sum
```go
import (
	"math"
)
```

```go
func Sum(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
```

### SumF
```go
import (
	"math"
)
```

```go
func SumF(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}
```

### Add
```go
import (
	"math"
)
```

```go
func Add(num1, num2 int) int {
	return num1 + num2
}
```

### AddF
```go
import (
	"math"
)
```

```go
func AddF(num1, num2 float64) float64 {
	return num1 + num2
}
```

### Sub
```go
import (
	"math"
)
```

```go
func Sub(num1, num2 int) int {
	return num1 - num2
}
```

### SubF
```go
import (
	"math"
)
```

```go
func SubF(num1, num2 float64) float64 {
	return num1 - num2
}
```

### Multi
```go
import (
	"math"
)
```

```go
func Multi(nums ...int) int {
	product := 1
	for _, v := range nums {
		product *= v
	}
	return product
}
```

### MultiF
```go
import (
	"math"
)
```

```go
func MultiF(nums ...float64) float64 {
	product := 1.0
	for _, v := range nums {
		product *= v
	}
	return product
}
```

### Divide
```go
import (
	"math"
)
```

```go
func Divide(num1, num2 int) int {
	return num1 / num2
}
```

### DivideF
```go
import (
	"math"
)
```

```go
func DivideF(num1, num2 float64) float64 {
	return num1 / num2
}
```

### Remainder
```go
import (
	"math"
)
```

```go
func Remainder(num1, num2 int) int {
	return num1 % num2
}
```

### Mean
```go
import (
	"math"
)
```

```go
func Mean(nums ...int) int {
	return Sum(nums...) / len(nums)
}
```

### MeanF
```go
import (
	"math"
)
```

```go
func MeanF(nums ...float64) float64 {
	return SumF(nums...) / float64(len(nums))
}
```

### Pow
```go
import (
	"math"
)
```

```go
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
```

### PowF
```go
import (
	"math"
)
```

```go
func PowF(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}
```

### Pow2
```go
import (
	"math"
)
```

```go
func Pow2(base int) int {
	return Multi(base, base)
}
```

### Pow2F
```go
import (
	"math"
)
```

```go
func Pow2F(base float64) float64 {
	return MultiF(base, base)
}
```

### Pow3
```go
import (
	"math"
)
```

```go
func Pow3(base int) int {
	return Multi(base, base, base)
}
```

### Pow3F
```go
import (
	"math"
)
```

```go
func Pow3F(base float64) float64 {
	return MultiF(base, base, base)
}
```

### Factorial
```go
import (
	"math"
)
```

```go
func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}
```

### Root
```go
import (
	"math"
)
```

```go
func Root(number, n int) int {
	return int(math.Pow(float64(number), 1.0/float64(n)))
}
```

### RootF
```go
import (
	"math"
)
```

```go
func RootF(number, n float64) float64 {
	return math.Pow(number, 1.0/n)
}
```

### Sqrt
```go
import (
	"math"
)
```

```go
func Sqrt(number int) int {
	return int(math.Sqrt(float64(number)))
}
```

### SqrtF
```go
import (
	"math"
)
```

```go
func SqrtF(number float64) float64 {
	return math.Sqrt(number)
}
```

### Cbrt
```go
import (
	"math"
)
```

```go
func Cbrt(number int) int {
	return Root(number, 3)
}
```

### CbrtF
```go
import (
	"math"
)
```

```go
func CbrtF(number float64) float64 {
	return RootF(number, 3)
}
```
