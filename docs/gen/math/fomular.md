# Fomular


### IsPrime

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
func IsEven(num int) bool {
	return num%2 == 0
}
```

### IsOdd

```go
func IsOdd(num int) bool {
	return num%2 != 0
}
```

### IsPerfectSquare

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
func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}
```

### LogF

```go
func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}
```

### Log2

```go
func Log2(number int) int {
	return Log(number, 2)
}
```

### Log2F

```go
func Log2F(number float64) float64 {
	return LogF(number, 2)
}
```

### Log10

```go
func Log10(number int) int {
	return Log(number, 10)
}
```

### Log10F

```go
func Log10F(number float64) float64 {
	return LogF(number, 10)
}
```

### Pytago

```go
func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}
```

### PytagoF

```go
func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}
```

### Sin

```go
func Sin(angle int) float64 {
	return math.Sin(float64(angle))
}
```

### SinF

```go
func SinF(angle float64) float64 {
	return math.Sin(angle)
}
```

### Cos

```go
func Cos(angle int) float64 {
	return math.Cos(float64(angle))
}
```

### CosF

```go
func CosF(angle float64) float64 {
	return math.Cos(angle)
}
```

### Tan

```go
func Tan(angle int) float64 {
	return math.Tan(float64(angle))
}
```

### TanF

```go
func TanF(angle float64) float64 {
	return math.Tan(angle)
}
```

### Speed

```go
func Speed(distance, time float64) float64 {
	return DivideF(distance, time)
}
```

### Distance

```go
func Distance(speed, time float64) float64 {
	return MultiF(speed, time)
}
```

### Time

```go
func Time(distance, speed float64) float64 {
	return DivideF(distance, speed)
}
```
