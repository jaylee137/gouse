# Write

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/io"
)
```
## Functions


### SampleIoWritePath

```go
func SampleIoWritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := io.WritePath(relativePath, newContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}
```
