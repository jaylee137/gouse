# Check


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
