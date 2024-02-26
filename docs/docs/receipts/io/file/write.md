# Write

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/io"
)
```
## Functions


### SampleIoWriteToFile

```go
func SampleIoWriteToFile() {
	err := io.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}
```
