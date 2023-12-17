# Time

## Imports

```go
import (
	"time"
)
```
## Functions


### format

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
func Second() int {
	return time.Now().Second()
}
```

### Minute

```go
func Minute() int {
	return time.Now().Minute()
}
```

### Hour

```go
func Hour() int {
	return time.Now().Hour()
}
```

### Day

```go
func Day() int {
	return time.Now().Day()
}
```

### Month

```go
func Month() int {
	return int(time.Now().Month())
}
```

### Year

```go
func Year() int {
	return time.Now().Year()
}
```

### Weekday

```go
func Weekday() int {
	return int(time.Now().Weekday())
}
```

### Unix

```go
func Unix() int64 {
	return time.Now().Unix()
}
```

### UnixMilli

```go
func UnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
```

### UnixMicro

```go
func UnixMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}
```

### UnixNano

```go
func UnixNano() int64 {
	return time.Now().UnixNano()
}
```

### UnixMilliToTime

```go
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(0, milli*int64(time.Millisecond))
}
```

### UnixMicroToTime

```go
func UnixMicroToTime(micro int64) time.Time {
	return time.Unix(0, micro*int64(time.Microsecond))
}
```

### UnixNanoToTime

```go
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(0, nano)
}
```
