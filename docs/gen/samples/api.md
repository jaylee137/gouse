# Api

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong1009/gouse/api")
```
## Functions


### apiPortScanner

```go
func apiPortScanner() {
	fmt.Println("Port Scanning")
	// <protocol>, <hostname>, <start port>, <end port>
	api.PortScanner("tcp", "127.0.0.1", 3000, 8080)
}```

### apiPortChecker

```go
func apiPortChecker() {
	fmt.Println("Port Checker")
	// <protocol>, <hostname>, <port>
	open := api.PortChecker("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}```
