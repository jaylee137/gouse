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

func Add(num1, num2 int) int {
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

func Divide(num1, num2 int) int {
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

func PowF(base, exp float64) int {
	return int(math.Pow(base, exp))
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