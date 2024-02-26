package array

import "github.com/thuongtruong1009/gouse/array"

func SampleArrayIncludes() {
	println("--- Check element is exist in array ---")
	println("[int]: ", array.Includes([]int{1, -2, 3}, 1))
	println("[uint]: ", array.Includes([]uint{1, 2, 3}, 1))
	println("[float]: ", array.Includes([]float64{1.2, 2.3, 3.4}, 1.2))
	println("[string]: ", array.Includes([]string{"1", "2", "3"}, "0"))
	println("[rune]: ", array.Includes([]rune{'a', 'b', 'c'}, 'a'))
	println("[complex]: ", array.Includes([]complex128{1 + 2i, 2 + 3i}, 1+2i))
	println("[struct]: ", array.Includes([]struct{ a int }{{1}, {2}}, struct{ a int }{3}))
}
