package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/math"
)

func mathIsPrime() {
	var num = 10
	println("Check prime number: ", math.IsPrime(num))
}

func mathIsEven() {
	var num = 10
	println("Check even number: ", math.IsEven(num))
}

func mathIsOdd() {
	var num = 10
	println("Check odd number: ", math.IsOdd(num))
}

func mathIsPerfectSquare() {
	var num = 10
	println("Check perfect square number: ", math.IsPerfectSquare(num))
}

func mathAbs() {
	var num = -10
	println("Absolute number: ", math.Abs(num))
}

func mathFloor() {
	var num = 10.49
	println("Floor number: ", math.Floor(num))
}

func mathCeil() {
	var num = 10.49
	println("Ceil number: ", math.Ceil(num))
}

func mathRound() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", math.Round(num1), math.Round(num2))
}

func mathMinMax() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Max number: ", math.Max(num1, num2, num3, num4))
	println("Min number: ", math.Min(num1, num2, num3, num4))
}

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

func mathPower() {
	println("Power of numbers: ", math.Pow(2, 3))
}

func mathFactorial() {
	var num = 5
	println("Factorial of number: ", math.Factorial(num))
}

func mathRoot() {
	println("Square-Root of integer number: ", math.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", math.SqrtF(20.0)))
	
	println("Cube-Root of integer number: ", math.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", math.CbrtF(20.0)))

	println("Nth-Root of integer number: ", math.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", math.RootF(20.0, 3.0)))
}

func mathLog() {
	println("Logarithm of integer number: ", math.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", math.LogF(20.0, 2.0)))
	
	println("Logarithm of integer number (base 2): ", math.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", math.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", math.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", math.Log10F(20.0)))
}

func MathExample() {
	mathIsPrime()
	mathIsEven()
	mathIsOdd()
	mathIsPerfectSquare()

	mathAbs()
	mathFloor()
	mathCeil()
	mathRound()
	mathMinMax()
	mathOperators()
	mathPower()
	mathFactorial()
	mathRoot()
	mathLog()
}