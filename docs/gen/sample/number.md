# Number

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/number"
)
```
## Functions


### numRandom

```go
func numRandom() {
	random := number.Random(1, 10)
	println("Random number [1, 10): ", random)
}
```

### numClamp

```go
func numClamp() {
	println("Clamp number: ", number.Clamp(5, 1, 10))
}
```

### numInRange

```go
func numInRange() {
	println("Check number is in range: ", number.InRange(5, 1, 10))
}
```

### NumberExample

```go
func NumberExample() {
	numRandom()
	numClamp()
	numInRange()
}
```
