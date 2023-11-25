package math

import "math"

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// down to the nearest integer
func Floor(num float64) int {
	return int(num)
}

// up to the nearest integer
func Ceil(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	if num > 0 {
		return int(num) + 1
	}
	return int(num) - 1
}

// round to the nearest integer
func Round(num float64) int {
	if num < 0 {
		return int(num)
	}

	if num - float64(int(num)) < 0.5 {
		return int(num)
	}

	return int(num) + 1
}

func Max(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(nums ...int) int {
	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func Sum(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumF(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func AddF(num1, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Multi(nums ...int) int {
	product := 1
	for _, v := range nums {
		product *= v
	}
	return product
}

func MultiF(nums ...float64) float64 {
	product := 1.0
	for _, v := range nums {
		product *= v
	}
	return product
}

func Divide(num1, num2 int) int {
	return num1 / num2
}

func DivideF(num1, num2 float64) float64 {
	return num1 / num2
}

func Remainder(num1, num2 int) int {
	return num1 % num2
}

func Mean(nums ...int) int {
	return Sum(nums...) / len(nums)
}

func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func Pow2(base int) int {
	return Multi(base, base)
}

func PowF(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}

func Factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * Factorial(num - 1)
}

func Root(number, n int) int {
	return int(math.Pow(float64(number), 1.0/float64(n)))
}

func RootF(number, n float64) float64 {
	return math.Pow(number, 1.0/n)
}

func Sqrt(number int) int {
	return int(math.Sqrt(float64(number)))
}

func SqrtF(number float64) float64 {
	return math.Sqrt(number)
}

func Cbrt(number int) int {
	return Root(number, 3)
}

func CbrtF(number float64) float64 {
	return RootF(number, 3)
}