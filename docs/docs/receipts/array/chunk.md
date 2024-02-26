# Chunk

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/array"
)
```
## Functions


### SampleArrayChunk

```go
func SampleArrayChunk() {
	println("--- Chunk array ---")
	fmt.Println("[int]: ", array.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	fmt.Println("[uint]: ", array.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("[float]: ", array.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	fmt.Println("[string]: ", array.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	fmt.Println("[rune]: ", array.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	fmt.Println("[complex]: ", array.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	fmt.Println("[struct]: ", array.Chunk([]struct{ a int }{{1}, {2}, {3}, {4}, {5}, {6}}, 3))
}
```
