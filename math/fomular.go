package math

import "math"

func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}

func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}

func Log2(number int) int {
	return Log(number, 2)
}

func Log2F(number float64) float64 {
	return LogF(number, 2)
}

func Log10(number int) int {
	return Log(number, 10)
}

func Log10F(number float64) float64 {
	return LogF(number, 10)
}

func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}

func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}