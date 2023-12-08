
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
		return format(date[0].(time.Time), "2006-01-02T15:04:05.999Z")
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
