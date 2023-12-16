# Ultilities


### format
```go
import (
	"time"
)
```

```go
func format(input interface{}, format string) string {
	var t time.Time

	switch v := input.(type) {
	case time.Time:
		t = v
	case string:
		var err error
		t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return ""
		}
	default:
		return ""
	}

	return t.Format(format)
}
```

### ISO
```go
import (
	"time"
)
```

```go
func ISO(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02T15:04:05.999Z")
	} else {
		return format(date[0], "2006-01-02T15:04:05.999Z")
	}
}
```

### Short
```go
import (
	"time"
)
```

```go
func Short(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "02/01/2006")
	} else {
		return format(date[0].(time.Time), "02/01/2006")
	}
}
```

### ShortNormal
```go
import (
	"time"
)
```

```go
func ShortNormal(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/02/2006")
	} else {
		return format(date[0].(time.Time), "01/02/2006")
	}
}
```

### ShortReverse
```go
import (
	"time"
)
```

```go
func ShortReverse(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006/01/02")
	} else {
		return format(date[0].(time.Time), "2006/01/02")
	}
}
```

### ShortDash
```go
import (
	"time"
)
```

```go
func ShortDash(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006-01-02")
	} else {
		return format(date[0].(time.Time), "2006-01-02")
	}
}
```

### ShortDot
```go
import (
	"time"
)
```

```go
func ShortDot(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006.01.02")
	} else {
		return format(date[0].(time.Time), "2006.01.02")
	}
}
```

### ShortUnderscore
```go
import (
	"time"
)
```

```go
func ShortUnderscore(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006_01_02")
	} else {
		return format(date[0].(time.Time), "2006_01_02")
	}
}
```

### ShortSpace
```go
import (
	"time"
)
```

```go
func ShortSpace(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "2006 01 02")
	} else {
		return format(date[0].(time.Time), "2006 01 02")
	}
}
```

### ShortMonth
```go
import (
	"time"
)
```

```go
func ShortMonth(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "01/2006")
	} else {
		return format(date[0].(time.Time), "01/2006")
	}
}
```

### Long
```go
import (
	"time"
)
```

```go
func Long(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Monday, January 2, 2006")
	} else {
		return format(date[0].(time.Time), "Monday, January 2, 2006")
	}
}
```

### UTC
```go
import (
	"time"
)
```

```go
func UTC(date ...interface{}) string {
	if len(date) == 0 {
		return format(time.Now(), "Jan 2, 2006 at 3:04pm (MST)")
	} else {
		return format(date[0].(time.Time), "Jan 2, 2006 at 3:04pm (MST)")
	}
}
```

### Second
```go
import (
	"time"
)
```

```go
func Second() int {
	return time.Now().Second()
}
```

### Minute
```go
import (
	"time"
)
```

```go
func Minute() int {
	return time.Now().Minute()
}
```

### Hour
```go
import (
	"time"
)
```

```go
func Hour() int {
	return time.Now().Hour()
}
```

### Day
```go
import (
	"time"
)
```

```go
func Day() int {
	return time.Now().Day()
}
```

### Month
```go
import (
	"time"
)
```

```go
func Month() int {
	return int(time.Now().Month())
}
```

### Year
```go
import (
	"time"
)
```

```go
func Year() int {
	return time.Now().Year()
}
```

### Weekday
```go
import (
	"time"
)
```

```go
func Weekday() int {
	return int(time.Now().Weekday())
}
```

### Unix
```go
import (
	"time"
)
```

```go
func Unix() int64 {
	return time.Now().Unix()
}
```

### UnixMilli
```go
import (
	"time"
)
```

```go
func UnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
```

### UnixMicro
```go
import (
	"time"
)
```

```go
func UnixMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}
```

### UnixNano
```go
import (
	"time"
)
```

```go
func UnixNano() int64 {
	return time.Now().UnixNano()
}
```

### UnixMilliToTime
```go
import (
	"time"
)
```

```go
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}
```

### UnixMicroToTime
```go
import (
	"time"
)
```

```go
func UnixMicroToTime(micro int64) time.Time {
	return time.Unix(0, micro*int64(time.Microsecond))
}
```

### UnixNanoToTime
```go
import (
	"time"
)
```

```go
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(0, nano)
}
```

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
