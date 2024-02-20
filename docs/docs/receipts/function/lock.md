# Lock

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong1009/gouse/function")
```
## Functions


### funcLock

```go
func funcLock() {
	oneInOneOutFuc := function.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	fmt.Println(oneInOneOutFuc)

	twoInTwoOutFunc1, twoInTwoOutFunc2 := function.LockFunc(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	}).(func(interface{}, interface{}) (interface{}, interface{}))("two input - two output (a)", "two input - two output (b)")
	fmt.Println(twoInTwoOutFunc1, twoInTwoOutFunc2)

	function.LockFunc(func() {
		println("no input - no output")
	}).(func())()
}```
