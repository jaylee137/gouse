# Max

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/array"
)
```
## Functions


### SampleArrayMax

```go
func SampleArrayMax() {
	println("--- Max element in array ---")
	println("[int]: ", array.Max([]int{1, -2, 3}))
	println("[uint]: ", array.Max([]uint{1, 2, 3}))
	println("[string]: ", array.Max([]string{"z", "d", "m"}))
	println("[rune]: ", string(array.Max([]rune{'z', 'd', 'm'})))
	println("[float]: ", array.Max([]float64{1.2, 2.3, 3.4}))
}
```
