# Tmp

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


### SampleCacheTmp

```go
func SampleCacheTmp() {
	newCache := cache.NewTmp(date.ToSecond(3))
	newCache.SetTmp("key1", "temp cache value 1", date.ToSecond(3))
	newCache.SetTmp("key2", "temp cache value 2", date.ToSecond(6))
	newCache.SetTmp("key3", "temp cache value 3", date.ToSecond(6))

	all := newCache.AllTmp()
	println("All temp cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	getted := newCache.GetTmp("key1")
	fmt.Println("Getted key 1 (before expires):", getted)

	time.Sleep(date.ToSecond(4))

	getted = newCache.GetTmp("key1")
	fmt.Println("Getted key 1 (after expires):", getted)

	newCache.DelTmp("key2")

	all = newCache.AllTmp()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	newCache.FlushTmp()

	all = newCache.AllTmp()
	fmt.Println("All temp cache values (after flush):", all)
}
```
