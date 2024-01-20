# Helper

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/helper"
)
```
## Functions


### helperRandomID

```go
func helperRandomID() {
	println("Generate random ID: ", helper.RandomID())
}
```

### helperUUID

```go
func helperUUID() {
	println("New uuid: ", helper.UUID())
}
```

### helperAutoMdDoc

```go
func helperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}
```
