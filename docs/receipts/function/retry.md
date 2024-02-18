# Retry

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/function")
```
## Functions


### funcRetry

```go
func funcRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}```
