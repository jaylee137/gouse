package array

func Most[T comparable](arr []T) T {
	var most = make(map[T]int)
	max := arr[0]
	for _, v := range arr {
		most[v] = most[v] + 1
		if(most[max] < most[v]) {
			max = v
		}
	}
	return max
}