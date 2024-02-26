# Info

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/io"
)
```
## Functions


### SampleIoFileInfo

```go
func SampleIoFileInfo() {
	data, err := io.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("File info: %+v\n", data.All)
	fmt.Println("File info (with name):", data.Name)
	fmt.Printf("File info (with size): %d bytes\n", data.Size)
	fmt.Println("File info (with permissions):", data.Mode)
	fmt.Println("File info (with last modified):", data.ModTime)
	fmt.Println("File info (with directory check): ", data.IsDir)
	fmt.Printf("File info (with system process): %+v\n", data.Sys)
}
```
