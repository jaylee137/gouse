# Date

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/date"
)
```
## Functions


### dateTime

```go
func dateTime() {
	println("Second:", date.Second())
	println("Minute:", date.Minute())
	println("Hour:", date.Hour())
	println("Day:", date.Day())
	println("Month:", date.Month())
	println("Year:", date.Year())
	println("Weekday:", date.Weekday())
	println("Unix:", date.Unix())
	println("UnixNano:", date.UnixNano())
	println("UnixMilli:", date.UnixMilli())
	println("UnixMicro:", date.UnixMicro())
	fmt.Println("UnixMilliToTime:", date.UnixMilliToTime(1000000000))
	fmt.Println("UnixMicroToTime:", date.UnixMicroToTime(1000000000))
	fmt.Println("UnixNanoToTime:", date.UnixNanoToTime(1000000000))
}
```

### dateISO

```go
func dateISO() {
	println("ISO:", date.ISO())
}
```

### dateShort

```go
func dateShort() {
	println("Short:", date.Short())
	println("ShortNormal:", date.ShortNormal())
	println("ShortReverse:", date.ShortReverse())
	println("ShortDash:", date.ShortDash())
	println("ShortDot:", date.ShortDot())
	println("ShortUnderscore:", date.ShortUnderscore())
	println("ShortSpace:", date.ShortSpace())
	println("ShortMonth:", date.ShortMonth())
}
```

### dateLong

```go
func dateLong() {
	println("Long:", date.Long())
}
```

### dateUTC

```go
func dateUTC() {
	println("UTC:", date.UTC())
}
```

### dateToSecond

```go
func dateToSecond() {
	println("ToSecond:", date.ToSecond(1))
}
```

### dateToMinute

```go
func dateToMinute() {
	println("ToMinute:", date.ToMinute(1))
}
```

### dateToHour

```go
func dateToHour() {
	println("ToHour:", date.ToHour(1))
}
```

### dateSleepSecond

```go
func dateSleepSecond() {
	date.SleepSecond(1)
}
```

### dateSleepMinute

```go
func dateSleepMinute() {
	date.SleepMinute(1)
}
```

### dateSleepHour

```go
func dateSleepHour() {
	date.SleepHour(1)
}
```

### DateExample

```go
func DateExample() {
	dateTime()

	dateISO()
	dateShort()
	dateLong()
	dateUTC()

	dateToSecond()
	dateToMinute()
	dateToHour()
	dateSleepSecond()
	dateSleepMinute()
	dateSleepHour()
}
```
