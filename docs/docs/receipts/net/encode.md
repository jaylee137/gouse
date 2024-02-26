# Encode

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/net"
)
```
## Functions


### SampleNetEncode

```go
func SampleNetEncode() {
	println("Encode: ", net.Encode("https://google.com"))
}
```

### SampleNetDecode

```go
func SampleNetDecode() {
	println("Decode: ", net.Decode("https%3A%2F%2Fgoogle.com"))
}
```
