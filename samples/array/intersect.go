package array

import (
	"github.com/thuongtruong1009/gouse/array"
)

func SampleArrayIntersect() {
	println("--- Intersection arrays ---")
	println("[int]: ", array.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	println("[uint]: ", array.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", array.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", array.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", array.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", array.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", array.Intersect([]struct{ a int }{{1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
