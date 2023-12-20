# Math

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/math"
)
```
## Functions


### mathIsPrime

```go
func mathIsPrime() {
	var num = 10
	println("Check prime number: ", math.IsPrime(num))
}
```

### mathIsEven

```go
func mathIsEven() {
	var num = 10
	println("Check even number: ", math.IsEven(num))
}
```

### mathIsOdd

```go
func mathIsOdd() {
	var num = 10
	println("Check odd number: ", math.IsOdd(num))
}
```

### mathIsPerfectSquare

```go
func mathIsPerfectSquare() {
	var num = 10
	println("Check perfect square number: ", math.IsPerfectSquare(num))
}
```

### mathAbs

```go
func mathAbs() {
	var num = -10
	println("Absolute number: ", math.Abs(num))
}
```

### mathFloor

```go
func mathFloor() {
	var num = 10.49
	println("Floor number: ", math.Floor(num))
}
```

### mathCeil

```go
func mathCeil() {
	var num = 10.49
	println("Ceil number: ", math.Ceil(num))
}
```

### mathRound

```go
func mathRound() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", math.Round(num1), math.Round(num2))
}
```

### mathMinMax

```go
func mathMinMax() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Max number: ", math.Max(num1, num2, num3, num4))
	println("Min number: ", math.Min(num1, num2, num3, num4))
}
```

### mathOperators

```go
func mathOperators() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", math.Sum(num1, num2, num3, num4))
	println("Add numbers: ", math.Add(num1, num2))
	println("Subtract of numbers: ", math.Sub(num1, num2))
	println("Multiply of numbers: ", math.Multi(num1, num2, num3, num4))
	println("Quotient of numbers: ", math.Divide(num1, num2))
	println("Remainder of numbers: ", math.Remainder(num1, num2))
	println("Average/Mean of numbers: ", math.Mean(num1, num2, num3, num4))
}
```

### mathPower

```go
func mathPower() {
	println("Power square of number: ", math.Pow2(2))
	println("Power of integer numbers: ", math.Pow(2, 3))
	println("Power of float numbers: ", math.PowF(2.0, 3.0))
}
```

### mathFactorial

```go
func mathFactorial() {
	var num = 5
	println("Factorial of number: ", math.Factorial(num))
}
```

### mathRoot

```go
func mathRoot() {
	println("Square-Root of integer number: ", math.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", math.SqrtF(20.0)))

	println("Cube-Root of integer number: ", math.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", math.CbrtF(20.0)))

	println("Nth-Root of integer number: ", math.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", math.RootF(20.0, 3.0)))
}
```

### mathLog

```go
func mathLog() {
	println("Logarithm of integer number: ", math.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", math.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", math.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", math.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", math.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", math.Log10F(20.0)))
}
```

### mathPytago

```go
func mathPytago() {
	println("Pytago of number (integer): ", fmt.Sprintf("%f", math.Pytago(3, 4)))
	println("Pytago of number (float): ", fmt.Sprintf("%f", math.PytagoF(3.0, 4.0)))
}
```

### mathTrigonometry

```go
func mathTrigonometry() {
	println("Sine of integer number: ", math.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", math.SinF(90.0)))
	println("Cosine of integer number: ", math.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", math.CosF(90.0)))
	println("Tangent of integer number: ", math.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", math.TanF(90.0)))
}
```

### mathRect

```go
func mathRect() {
	println("Area of rectangle: ", math.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", math.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", math.PeriRectF(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%f", math.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", math.DiagRectF(10.0, 20.0)))
	println("Volume of rectangular (integer): ", math.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", math.VolRectF(10.0, 20.0, 30.0)))
}
```

### mathCircle

```go
func mathCircle() {
	println("Area of circle (integer): ", fmt.Sprintf("%f", math.AreaCircle(10)))
	println("Area of circle (float): ", fmt.Sprintf("%f", math.AreaCircleF(10.0)))
	println("Perimeter of circle (integer): ", fmt.Sprintf("%f", math.PeriCircle(10)))
	println("Perimeter of circle (float): ", fmt.Sprintf("%f", math.PeriCircleF(10.0)))
}
```

### mathTriangle

```go
func mathTriangle() {
	println("Area of triangle (integer): ", math.AreaTriangle(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", math.AreaTriangleF(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", math.PeriTriangle(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", math.PeriTriangleF(10.0, 20.0, 30.0)))
}
```

### mathSquare

```go
func mathSquare() {
	println("Area of square (integer): ", math.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", math.AreaSquareF(10.0)))
	println("Perimeter of square (integer): ", math.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", math.PeriSquareF(10.0)))
}
```

### mathCube

```go
func mathCube() {
	println("Area of cube (integer): ", math.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", math.AreaCubeF(10.0)))
	println("Perimeter of cube (integer): ", math.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", math.PeriCubeF(10.0)))
	println("Volume of cube (integer): ", math.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", math.VolCubeF(10.0)))
}
```

### mathSphere

```go
func mathSphere() {
	println("Area of sphere (integer): ", fmt.Sprintf("%f", math.AreaSphere(10)))
	println("Area of sphere (float): ", fmt.Sprintf("%f", math.AreaSphereF(10.0)))
	println("Volume of sphere (integer): ", fmt.Sprintf("%f", math.VolSphere(10)))
	println("Volume of sphere (float): ", fmt.Sprintf("%f", math.VolSphereF(10.0)))
}
```

### mathCylinder

```go
func mathCylinder() {
	println("Area of cylinder (integer): ", fmt.Sprintf("%f", math.AreaCylinder(10, 20)))
	println("Area of cylinder (float): ", fmt.Sprintf("%f", math.AreaCylinderF(10.0, 20.0)))
	println("Volume of cylinder (integer): ", fmt.Sprintf("%f", math.VolCylinder(10, 20)))
	println("Volume of cylinder (float): ", fmt.Sprintf("%f", math.VolCylinderF(10.0, 20.0)))
}
```

### mathCone

```go
func mathCone() {
	println("Area of cone (integer): ", fmt.Sprintf("%f", math.AreaCone(10, 20)))
	println("Area of cone (float): ", fmt.Sprintf("%f", math.AreaConeF(10.0, 20.0)))
	println("Volume of cone (integer): ", fmt.Sprintf("%f", math.VolCone(10, 20)))
	println("Volume of cone (float): ", fmt.Sprintf("%f", math.VolConeF(10.0, 20.0)))
}
```

### mathTrapezoid

```go
func mathTrapezoid() {
	println("Area of trapezoid (integer): ", fmt.Sprintf("%f", math.AreaTrapezoid(10, 20, 30)))
	println("Area of trapezoid (float): ", fmt.Sprintf("%f", math.AreaTrapezoidF(10.0, 20.0, 30.0)))
}
```

### mathParallelogram

```go
func mathParallelogram() {
	println("Area of parallelogram (integer): ", math.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", math.AreaParallelogramF(10.0, 20.0)))
}
```

### mathRhombus

```go
func mathRhombus() {
	println("Area of rhombus (integer): ", math.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", math.AreaRhombusF(10.0, 20.0)))
}
```

### mathEllipse

```go
func mathEllipse() {
	println("Area of ellipse (integer): ", fmt.Sprintf("%f", math.AreaEllipse(10, 20)))
	println("Area of ellipse (float): ", fmt.Sprintf("%f", math.AreaEllipseF(10.0, 20.0)))
}
```

### mathPolygon

```go
func mathPolygon() {
	println("Area of pentagol by number of sides (integer): ", fmt.Sprintf("%f", math.AreaPolygon(10, 6)))
}
```

### mathTransition

```go
func mathTransition() {
	var distance, speed, time float64 = 100, 10, 10
	println("Speed: ", fmt.Sprintf("%f", math.Speed(distance, time)))
	println("Distance: ", fmt.Sprintf("%f", math.Distance(speed, time)))
	println("Time: ", fmt.Sprintf("%f", math.Time(distance, speed)))
}
```
