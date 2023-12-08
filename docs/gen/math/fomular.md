
### IsPrime
```go
import ()
```

```go
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
```

### IsEven
```go
import ()
```

```go
func IsEven(num int) bool {
	return num%2 == 0
}
```

### IsOdd
```go
import ()
```

```go
func IsOdd(num int) bool {
	return num%2 != 0
}
```

### IsPerfectSquare
```go
import ()
```

```go
func IsPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
```

### Log
```go
import (
	"math"
)
```

```go
func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}
```

### LogF
```go
import (
	"math"
)
```

```go
func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}
```

### Log2
```go
import (
	"math"
)
```

```go
func Log2(number int) int {
	return Log(number, 2)
}
```

### Log2F
```go
import (
	"math"
)
```

```go
func Log2F(number float64) float64 {
	return LogF(number, 2)
}
```

### Log10
```go
import (
	"math"
)
```

```go
func Log10(number int) int {
	return Log(number, 10)
}
```

### Log10F
```go
import (
	"math"
)
```

```go
func Log10F(number float64) float64 {
	return LogF(number, 10)
}
```

### Pytago
```go
import (
	"math"
)
```

```go
func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}
```

### PytagoF
```go
import (
	"math"
)
```

```go
func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}
```

### Sin
```go
import (
	"math"
)
```

```go
func Sin(angle int) float64 {
	return math.Sin(float64(angle))
}
```

### SinF
```go
import (
	"math"
)
```

```go
func SinF(angle float64) float64 {
	return math.Sin(angle)
}
```

### Cos
```go
import (
	"math"
)
```

```go
func Cos(angle int) float64 {
	return math.Cos(float64(angle))
}
```

### CosF
```go
import (
	"math"
)
```

```go
func CosF(angle float64) float64 {
	return math.Cos(angle)
}
```

### Tan
```go
import (
	"math"
)
```

```go
func Tan(angle int) float64 {
	return math.Tan(float64(angle))
}
```

### TanF
```go
import (
	"math"
)
```

```go
func TanF(angle float64) float64 {
	return math.Tan(angle)
}
```

### Speed
```go
import (
	"math"
)
```

```go
func Speed(distance, time float64) float64 {
	return DivideF(distance, time)
}
```

### Distance
```go
import (
	"math"
)
```

```go
func Distance(speed, time float64) float64 {
	return MultiF(speed, time)
}
```

### Time
```go
import (
	"math"
)
```

```go
func Time(distance, speed float64) float64 {
	return DivideF(distance, speed)
}
```
