# Read_by_line

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/io"
)
```
## Functions


### SampleIoReadFileByLine

```go
func SampleIoReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
```
