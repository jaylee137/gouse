# Cache

## Imports

```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong1009/gouse/cache"
	"github.com/thuongtruong1009/gouse/date"
)
```
## Functions


### cacheCacheLocal

```go
func cacheCacheLocal() {
	cache.SetLocal("key1", "local cache value 1")
	cache.SetLocal("key2", "local cache value 2")
	cache.SetLocal("key3", "local cache value 3")

	all := cache.AllLocal()
	println("All local cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := cache.GetLocal("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Getted key 1:", getted1)

	cache.DelLocal("key2")

	all = cache.AllLocal()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	cache.FlushLocal()

	all = cache.AllLocal()
	fmt.Println("All local cache values (after flush):", all)
}
```

### cacheCacheTmp

```go
func cacheCacheTmp() {
	cache.SetTmp("key1", "temp cache value 1", date.ToSecond(3))
	cache.SetTmp("key2", "temp cache value 2", date.ToSecond(6))
	cache.SetTmp("key3", "temp cache value 3", date.ToSecond(6))

	all := cache.AllTmp()
	println("All temp cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	getted := cache.GetTmp("key1")
	fmt.Println("Getted key 1 (before expires):", getted)

	time.Sleep(date.ToSecond(4))

	getted = cache.GetTmp("key1")
	fmt.Println("Getted key 1 (after expires):", getted)

	cache.DelTmp("key2")

	all = cache.AllTmp()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	cache.FlushTmp()

	all = cache.AllTmp()
	fmt.Println("All temp cache values (after flush):", all)
}
```
