# For_by

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong1009/gouse/array")
```
## Functions


### arrForBy

```go
func arrForBy() {
	println("--- Loop array then handler with callback function ---")
	print("[int]: ")
	array.ForBy([]int{1, -2, 3, -4, 5, 6}, func(v int) {
		println(v)
	})

	print("[uint]: ")
	array.ForBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) {
		println(v)
	})

	print("[float]: ")
	array.ForBy([]float64{1.2, 2.3, 3.4}, func(v float64) {
		println(v)
	})

	print("[string]: ")
	array.ForBy([]string{"1", "4", "5", "6"}, func(v string) {
		println(v)
	})

	print("[rune]: ")
	array.ForBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) {
		println(v)
	})

	print("[complex]: ")
	array.ForBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) {
		println(v)
	})

	print("[struct]: ")
	array.ForBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) {
		fmt.Println(v)
	})
}```
