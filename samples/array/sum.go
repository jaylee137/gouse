package array

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/array"
)

func SampleArraySum() {
	println("--- Sum elements in array ---")
	println("[int]: ", array.Sum([]int{1, -2, 3}))
	println("[uint]: ", array.Sum([]uint{1, 2, 3}))
	fmt.Println("[float]: ", array.Sum([]float64{1.2, 2.3, 3.4}))
	println("[rune]: ", array.Sum([]rune{'a', 'b', 'c'}))
	fmt.Println("[complex]: ", array.Sum([]complex128{1 + 2i, 2 + 3i, 3 + 4i}))
}
