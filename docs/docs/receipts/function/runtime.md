# Runtime

## Imports

```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong1009/gouse/function"
)
```
## Functions


### SampleFuncRunTime

```go
func SampleFuncRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := function.RunTime(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
```
