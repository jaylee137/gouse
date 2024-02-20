# Port_checker

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong1009/gouse/api")
```
## Functions


### SampleApiPortChecker

```go
func SampleApiPortChecker() {
	open := api.PortChecker("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}```
