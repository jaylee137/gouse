# Delay

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/function")
```
## Functions


### funcDelay

```go
func funcDelay() {
	println("Delay start:")

	result := function.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn {
		println(result.Value)
	} else {
		println("No result")
	}

	function.Delay(func() {
		println("Delayed not return")
	}, 3)
}```
