
### Random
```go
import (
	"math/rand"

	"time"
)
```

```go
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)

	return min + rand.Intn(max-min)
}
```

### Clamp
```go
import (
	"math/rand"

	"time"
)
```

```go
func Clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}
```

### InRange
```go
import (
	"math/rand"

	"time"
)
```

```go
func InRange(value int, min int, max int) bool {
	return value >= min && value <= max
}
```
