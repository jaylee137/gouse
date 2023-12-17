# Ultilities

## Imports

```go
import (
	"crypto/rand"

	"math/big"
)
```
## Functions


### Random

```go
func Random(min, max int) int {
	// [min, max)
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0
	}

	return int(randomInt.Int64()) + min
}
```

### Clamp

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
func InRange(value int, min int, max int) bool {
	return value >= min && value <= max
}
```
