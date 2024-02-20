# Equal

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/array")
```
## Functions


### arrEqual

```go
func arrEqual() {
	println("--- Compare equal ---")
	println("[int]: ", array.Equal(1, 1))
	println("[uint]: ", array.Equal(uint(1), uint(1)))
	println("[float]: ", array.Equal(1.2, 1.1))
	println("[string]: ", array.Equal("1", "0"))
	println("[rune]: ", array.Equal('a', 'a'))
	println("[bool]: ", array.Equal(true, true))
	println("[complex]: ", array.Equal(1+2i, 1+2i))
	println("[struct]: ", array.Equal(struct{ a int }{1}, struct{ a int }{1}))
}```
