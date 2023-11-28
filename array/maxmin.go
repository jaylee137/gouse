package array

func minmax[T comparable](arr []T, less func(T, T) bool) T {
	if len(arr) == 0 {
		panic("Empty array")
	}

	max := arr[0]
	for _, item := range arr {
		if less(max, item) {
			max = item
		}
	}
	return max
}

func Min[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return minmax[T](arr, func(a, b T) bool {
		return a > b
	})
}

func Max[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return minmax[T](arr, func(a, b T) bool {
		return a < b
	})
}
