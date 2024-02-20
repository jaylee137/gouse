# Index

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/array")
```
## Functions


### arrIndex

```go
func arrIndex() {
	println("--- Index of element in array ---")
	println("[int]: ", array.Index([]int{1, -2, 3, -4, 5, 6}, 3))
	println("[uint]: ", array.Index([]uint{1, 2, 3, 4, 5, 7}, 3))
	println("[float]: ", array.Index([]float64{1.2, 2.3, 3.4}, 3.4))
	println("[string]: ", array.Index([]string{"1", "4", "5", "6"}, "5"))
	println("[rune]: ", array.Index([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	println("[complex]: ", array.Index([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5+6i))
	println("[struct]: ", array.Index([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, struct{ a int }{3}))
}```
