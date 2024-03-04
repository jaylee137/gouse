# Date

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/date"
)
```
## Functions


### SampleDateTime

```go
func SampleDateTime() {
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

### SampleDateISO

```go
func SampleDateISO() {
	println("ISO:", date.ISO())
}
```

### SampleDateShort

```go
func SampleDateShort() {
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

### SampleDateLong

```go
func SampleDateLong() {
	println("Long:", date.Long())
}
```

### SampleDateUTC

```go
func SampleDateUTC() {
	println("UTC:", date.UTC())
}
```

### SampleDateToSecond

```go
func SampleDateToSecond() {
	println("ToSecond:", date.ToSecond(1))
}
```

### SampleDateToMinute

```go
func SampleDateToMinute() {
	println("ToMinute:", date.ToMinute(1))
}
```

### SampleDateToHour

```go
func SampleDateToHour() {
	println("ToHour:", date.ToHour(1))
}
```

### SampleDateSleepSecond

```go
func SampleDateSleepSecond() {
	date.SleepSecond(1)
}
```

### SampleDateSleepMinute

```go
func SampleDateSleepMinute() {
	date.SleepMinute(1)
}
```

### SampleDateSleepHour

```go
func SampleDateSleepHour() {
	date.SleepHour(1)
}
```

### SampleDateClock

```go
func SampleDateClock() {
	date.TerminalClock()
}
```
