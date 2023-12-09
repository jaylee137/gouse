# Utilities


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
