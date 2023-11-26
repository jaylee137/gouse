package math

import "math"

func AreaRect(length, width int) int {
	return Multi(length, width)
}

func AreaRectF(length, width float64) float64 {
	return MultiF(length, width)
}

func PeriRect(length, width int) int {
	return Multi(2, Add(length, width))
}

func PeriRectF(length, width float64) float64 {
	return MultiF(2, AddF(length, width))
}

func DiagRect(length, width int) float64 {
	return Pytago(length, width)
}

func DiagRectF(length, width float64) float64 {
	return PytagoF(length, width)
}

func VolRect(length, width, height int) int {
	return Multi(length, width, height)
}

func VolRectF(length, width, height float64) float64 {
	return MultiF(length, width, height)
}

func AreaCircle(radius int) float64 {
	return math.Pi * float64(Pow2(radius))
}

func AreaCircleF(radius float64) float64 {
	return math.Pi * PowF(radius, 2)
}

func PeriCircle(radius int) float64 {
	return 2 * math.Pi * float64(radius)
}

func PeriCircleF(radius float64) float64 {
	return 2 * math.Pi * radius
}

func AreaTriangle(base, height int)  int {
	return Divide(Multi(base, height), 2)
}

func AreaTriangleF(base, height float64) float64 {
	return DivideF(MultiF(base, height), 2)
}

func PeriTriangle(side1, side2, side3 int) int {
	return Sum(side1, side2, side3)
}

func PeriTriangleF(side1, side2, side3 float64) float64 {
	return SumF(side1, side2, side3)
}

func AreaSquare(side int) int {
	return Pow2(side)
}

func AreaSquareF(side float64) float64 {
	return PowF(side, 2)
}

func PeriSquare(side int) int {
	return Multi(side, 4)
}

func PeriSquareF(side float64) float64 {
	return MultiF(side, 4)
}

func AreaCube(side int) int {
	return Multi(Pow2(side), 6)
}

func AreaCubeF(side float64) float64 {
	return MultiF(PowF(side, 2), 6)
}

func PeriCube(side int) int {
	return Multi(side, 12)
}

func PeriCubeF(side float64) float64 {
	return MultiF(side, 12)
}

func VolCube(side int) int {
	return Pow3(side)
}

func VolCubeF(side float64) float64 {
	return PowF(side, 3)
}

func AreaSphere(radius int) float64 {
	return 4 * math.Pi * float64(Pow2(radius))
}

func AreaSphereF(radius float64) float64 {
	return 4 * math.Pi * PowF(radius, 2)
}

func VolSphere(radius int) float64 {
	return (4 / 3) * math.Pi * float64(Pow2(radius))
}

func VolSphereF(radius float64) float64 {
	return (4 / 3) * math.Pi * PowF(radius, 2)
}

func AreaCylinder(radius, height int) float64 {
	return (2 * math.Pi * float64(radius)) * float64(height)
}

func AreaCylinderF(radius, height float64) float64 {
	return (2 * math.Pi * radius) * height
}

func VolCylinder(radius, height int) float64 {
	return math.Pi * float64(Pow2(radius)) * float64(height)
}

func VolCylinderF(radius, height float64) float64 {
	return math.Pi * PowF(radius, 2) * height
}

func AreaCone(radius, height int) float64 {
	return math.Pi * float64(radius) * (float64(radius) + math.Sqrt(float64(Pow2(height) + Pow2(radius))))
}

func AreaConeF(radius, height float64) float64 {
	return math.Pi * radius * (radius + math.Sqrt(PowF(height, 2) + PowF(radius, 2)))
}

func VolCone(radius, height int) float64 {
	return DivideF(1, 3) * math.Pi * float64(Pow2(radius)) * float64(height)
}

func VolConeF(radius, height float64) float64 {
	return DivideF(1, 3) * math.Pi * PowF(radius, 2) * height
}

func AreaTrapezoid(base1, base2, height int) float64 {
	return 0.5 * float64(Add(base1, base2)) * float64(height)
}

func AreaTrapezoidF(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

func AreaParallelogram(base, height int) int {
	return Multi(base, height)
}

func AreaParallelogramF(base, height float64) float64 {
	return MultiF(base, height)
}

func AreaRhombus(diag1, diag2 int) int {
	return Divide(Multi(diag1, diag2), 2)
}

func AreaRhombusF(diag1, diag2 float64) float64 {
	return DivideF(MultiF(diag1, diag2), 2)
}

func AreaEllipse(major, minor int) float64 {
	return math.Pi * float64(major) * float64(minor)
}

func AreaEllipseF(major, minor float64) float64 {
	return math.Pi * major * minor
}

func AreaPolygon(lenSide float64, numSide int) float64 {
	return DivideF(MultiF(0.25, float64(numSide), Pow2F(lenSide)), AbsF(TanF(DivideF(180, float64(numSide)))))
}