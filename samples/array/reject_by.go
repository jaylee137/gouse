package array

import (
	"github.com/thuongtruong1009/gouse/array"
)

func arrRejectBy() {
	println("--- Reject elements in array by pass condition in callback function---")
	println("[int]: ", array.RejectBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v > 2
	}))

	println("[uint]: ", array.RejectBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v > 2
	}))

	println("[float]: ", array.RejectBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v > 2
	}))

	println("[string]: ", array.RejectBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v > "3"
	}))

	println("[rune]: ", array.RejectBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v > 'c'
	}))

	println("[complex]: ", array.RejectBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return real(v) > 3
	}))
}
