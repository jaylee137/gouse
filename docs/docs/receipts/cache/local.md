# Local

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/cache"
)
```
## Functions


### SampleCacheLocal

```go
func SampleCacheLocal() {
	newCache := cache.NewLocal()
	newCache.SetLocal("key1", "local cache value 1")
	newCache.SetLocal("key2", "local cache value 2")
	newCache.SetLocal("key3", "local cache value 3")

	all := newCache.AllLocal()
	println("All local cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := newCache.GetLocal("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Getted key 1:", getted1)

	newCache.DelLocal("key2")

	all = newCache.AllLocal()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	newCache.FlushLocal()

	all = newCache.AllLocal()
	fmt.Println("All local cache values (after flush):", all)
}
```
