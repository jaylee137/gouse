package array

import (
	"github.com/thuongtruong1009/gouse/array"
)

func SampleArrayMerge() {
	println("--- Merge arrays ---")
	println("[int]: ", array.Merge([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, []int{1, -2, 3, -4, 5, 6}))
	println("[uint]: ", array.Merge([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", array.Merge([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", array.Merge([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", array.Merge([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", array.Merge([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", array.Merge([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
