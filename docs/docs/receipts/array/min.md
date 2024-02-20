# Min

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/array")
```
## Functions


### arrMin

```go
func arrMin() {
	println("--- Min element in array ---")
	println("[int]: ", array.Min([]int{1, -2, 3}))
	println("[uint]: ", array.Min([]uint{1, 2, 3}))
	println("[string]: ", array.Min([]string{"z", "d", "m"}))
	println("[rune]: ", string(array.Min([]rune{'z', 'd', 'm'})))
	println("[float]: ", array.Min([]float64{1.2, 2.3, 3.4}))
}```
