# Ultilities


### ToSecond
```go
import (
	"time"
)
```

```go
func ToSecond(second int) time.Duration {
	return time.Duration(second) * time.Second
}
```

### ToMinute
```go
import (
	"time"
)
```

```go
func ToMinute(minute int) time.Duration {
	return time.Duration(minute) * time.Minute
}
```

### ToHour
```go
import (
	"time"
)
```

```go
func ToHour(hour int) time.Duration {
	return time.Duration(hour) * time.Hour
}
```

### SleepSecond
```go
import (
	"time"
)
```

```go
func SleepSecond(second int) {
	time.Sleep(ToSecond(second))
}
```

### SleepMinute
```go
import (
	"time"
)
```

```go
func SleepMinute(minute int) {
	time.Sleep(ToMinute(minute))
}
```

### SleepHour
```go
import (
	"time"
)
```

```go
func SleepHour(hour int) {
	time.Sleep(ToHour(hour))
}
```
