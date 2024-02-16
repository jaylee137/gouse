package array

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/array"
)

func arrDrop() {
	println("--- Drop elements in array (default n = 1) ---")
	fmt.Println("[int] with default: ", array.Drop([]int{1, -2, 3, -4, 5, 6}))
	fmt.Println("[int]: ", array.Drop([]int{1, -2, 3, -4, 5, 6}, 2))
	fmt.Println("[uint]: ", array.Drop([]uint{1, 2, 3, 4, 5, 7}, 2))
	fmt.Println("[float]: ", array.Drop([]float64{1.2, 2.3, 3.4}, 2))
	fmt.Println("[string]: ", array.Drop([]string{"1", "4", "5", "6"}, 2))
	fmt.Println("[rune]: ", array.Drop([]rune{'a', 'b', 'd', 'e', 'f'}, 2))
	fmt.Println("[complex]: ", array.Drop([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 2))
	fmt.Println("[struct]: ", array.Drop([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, 2))
}
