
### delay
```go
import (
	"time"
)
```

```go
func delay[T any](f func() T, timeout int, hasReturn bool) DelayedResult[T] {
	resultChan := make(chan T, 1)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		result := f()
		if hasReturn {
			resultChan <- result
		}
	}()

	if hasReturn {
		return DelayedResult[T]{Value: <-resultChan, HasReturn: true}
	}

	return DelayedResult[T]{HasReturn: false}
}
```

### DelayF
```go
import (
	"time"
)
```

```go
func DelayF[T any](f func() T, timeout int) DelayedResult[T] {
	return delay(f, timeout, true)
}
```

### Delay
```go
import (
	"time"
)
```

```go
func Delay(f func(), timeout int) DelayedResult[struct{}] {
	return delay(func() struct{} {
		f()
		return struct{}{}
	}, timeout, false)
}
```

### lock
```go
import (
	"reflect"

	"sync"
)
```

```go
func (mw *MutexWrapper) lock() {
	mw.mutex.Lock()
}
```

### unLock
```go
import (
	"reflect"

	"sync"
)
```

```go
func (mw *MutexWrapper) unLock() {
	mw.mutex.Unlock()
}
```

### LockFunc
```go
import (
	"reflect"

	"sync"
)
```

```go
func LockFunc(callback interface{}) interface{} {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()

	callbackType := reflect.TypeOf(callback)
	if callbackType.Kind() != reflect.Func {
		panic("callback must be a function")
	}

	return reflect.MakeFunc(callbackType, func(params []reflect.Value) []reflect.Value {
		// Convert params to interface{} slice
		var resultInterfaces []interface{}
		for _, result := range reflect.ValueOf(callback).Call(params) {
			resultInterfaces = append(resultInterfaces, result.Interface())
		}

		// Return the results as reflect.Value
		var resultReflectValues []reflect.Value
		for _, result := range resultInterfaces {
			resultReflectValues = append(resultReflectValues, reflect.ValueOf(result))
		}

		return resultReflectValues
	}).Interface()
}
```

### Retry
```go
import (
	"time"

	"github.com/thuongtruong1009/gouse/date"
)
```

```go
func Retry(fn func() error, attempts int, sleep int) (err error) {
	retry := func() {
		if err = fn(); err != nil {
			return
		}
	}

	for i := attempts; i > 0; i-- {
		if i > 1 {
			retry()
			date.SleepSecond(sleep)
		} else if i == 1 {
			retry()
		} else {
			return nil
		}
	}
	return
}
```

### Times
```go
import (
	"time"

	"github.com/thuongtruong1009/gouse/date"
)
```

```go
func Times(fn func(), attempts int) {
	for i := 0; i < attempts; i++ {
		fn()
	}
}
```

### Interval
```go
import (
	"time"

	"github.com/thuongtruong1009/gouse/date"
)
```

```go
func Interval(fn func(), timeout int) {
	for {
		fn()
		date.SleepSecond(timeout)
	}
}
```

### RunTime
```go
import (
	"time"

	"github.com/thuongtruong1009/gouse/date"
)
```

```go
func RunTime(startTime time.Time, task func()) time.Duration {
	task()
	elapsedTime := float64(time.Since(startTime).Seconds() * 1000)
	return time.Duration(elapsedTime)
}
```
