package array

import (
	"github.com/thuongtruong1009/gouse/array"
)

func arrIndexBy() {
	println("--- Find index of element pass condition in callback function ---")
	println("[int]: ", array.IndexBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", array.IndexBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", array.IndexBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", array.IndexBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", array.IndexBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", array.IndexBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", array.IndexBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}
