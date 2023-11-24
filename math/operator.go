package math

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Ceil(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	return int(num) + 1
}

func Floor(num float64) int {
	return int(num)
}

func Round(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	if num > 0 {
		return int(num) + 1
	}
	return int(num) - 1
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
	var min int
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

func SumBy(nums []int, fn func(int) int) int {
	var sum int
	for _, v := range nums {
		sum += fn(v)
	}
	return sum
}

func Average(nums ...int) int {
	return Sum(nums...) / len(nums)
}

func AverageBy(nums []int, fn func(int) int) int {
	return SumBy(nums, fn) / len(nums)
}