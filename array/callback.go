package array

func CbIndex[T comparable](arr []T, f func(T) bool) int {
	for i, v := range arr {
		if f(v) {
			return i
		}
	}
	return -1
}

func CbKey[T comparable](arr []T, f func(T) bool) int {
	return CbIndex(arr, f)
}

func CbFilter[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func CbFind[T comparable](arr []T, f func(T) bool) T {
	for _, v := range arr {
		if f(v) {
			return v
		}
	}
	return arr[0]
}

func CbFor[T comparable](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func CbMap[T comparable, R comparable](arr []T, f func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}