# Compact

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/array"
)
```
## Functions


### SampleArrayCompact

```go
func SampleArrayCompact() {
	result := array.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	fmt.Println("Compact remove all falsy values: ", result)
}
```
