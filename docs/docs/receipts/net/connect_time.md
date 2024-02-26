# Connect_time

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/net"
)
```
## Functions


### SampleNetConnectTime

```go
func SampleNetConnectTime() {
	connectTime, err := net.ConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```
