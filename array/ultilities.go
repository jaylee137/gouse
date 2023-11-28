package array

import "github.com/thuongtruong1009/gouse/types"

func Includes[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func Sum[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Most[T comparable](arr []T) T {
	var most = make(map[T]int)
	max := arr[0]
	for _, v := range arr {
		most[v] = most[v] + 1
		if most[max] < most[v] {
			max = v
		}
	}
	return max
}

func Chunk[T comparable](array []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}

func Diff[T comparable](a, b []T) []T {
	var diff []T
	for _, v := range a {
		if !Includes(b, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

func Drop[T comparable](arr []T, n ...int) []T {
	if len(n) == 0 {
		n = append(n, 1)
	}

	return arr[n[0]:]
}

func Index[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func Merge[T comparable](arr ...[]T) []T {
	var merged []T
	for _, v := range arr {
		merged = append(merged, v...)
	}
	return merged
}

func Compact[T interface{}](arr []T) []T {
	var compact []T
	for _, v := range arr {
		if !types.IsZero(v) && !types.IsNil(v) && !types.IsEmpty(v) && !types.IsUndefined(v) && !types.IsBool(v) {
			compact = append(compact, v)
		}
	}
	return compact
}
