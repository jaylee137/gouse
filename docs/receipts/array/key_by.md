# Key_by

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/array")
```
## Functions


### arrKeyBy

```go
func arrKeyBy() {
	println("--- Find key of element pass condition in callback function ---")
	println("[int]: ", array.KeyBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", array.KeyBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", array.KeyBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", array.KeyBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", array.KeyBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", array.KeyBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", array.KeyBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}```
