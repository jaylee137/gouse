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
	var cpuprofile, memprofile = "cpu.pprof", "mem.pprof"
	tools.Profile(cpuprofile, memprofile)

	// run this command to test
	// go tool pprof <cpu.pprof_or_mem.pprof>

	// How to generate visual analysis images
	//	Open https://graphviz.gitlab.io/download/ and follow the prompts to download and install.
	// After the installation is complete, for Windows, add the bin folder of the Graphviz installation path after setting the environment variable path.

	// Update main.go
	// 	package main
	// import _ "net/http/pprof"
	// import "net/http"
	// func main() {
	//     go func() {
	//         _ = http.ListenAndServe("0.0.0.0:8081", nil)
	//     }()
	//     // your code
	// }

	// How to check the memory size of each module
	// go tool pprof -alloc_space -cum http://localhost:8081/debug/pprof/heap

	// View specific data list in web browser
	// http://localhost:8081/debug/pprof/heap?debug=1

	// 	Command line method to generate visual analysis images
	// 	go tool pprof http://localhost:8081/debug/pprof/goroutine

	//	View specific data list in web browser
	//
	// http://localhost:8081/debug/pprof/goroutine?debug=1
}
```
