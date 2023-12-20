# Tool

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/tools"
)
```
## Functions


### toolDoc

```go
func toolDoc() {
	tools.Doc("docs/gen")
}
```

### toolProfile

```go
func toolProfile() {
	tools.Profile()
	// run this command to test
	// go tool pprof -http=:8080 <cpu.prof_or_mem.prof>
}
```
