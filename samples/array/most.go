package array

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/array"
)

func arrMost() {
	println("--- Most frequency in array ---")
	println("[int]: ", array.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", array.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	fmt.Println("[float]: ", array.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", array.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(array.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	fmt.Println("[complex]: ", array.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	fmt.Println("[struct]: ", array.Most([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}
