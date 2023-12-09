# Check


### IsLetter
```go
import ()
```

```go
func IsLetter[T byte | rune](char T) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
```

### IsDigit
```go
import ()
```

```go
func IsDigit[T byte | rune](char T) bool {
	return char >= '0' && char <= '9'
}
```

### Includes
```go
import ()
```

```go
func Includes(s string, substr string) bool {
	return FIndex(s, substr) != -1
}
```

### StartsWith
```go
import ()
```

```go
func StartsWith(s string, substr string) bool {
	return FIndex(s, substr) == 0
}
```

### EndsWith
```go
import ()
```

```go
func EndsWith(s string, substr string) bool {
	return LIndex(s, substr) == len(s)-len(substr)
}
```

### IsLower
```go
import ()
```

```go
func IsLower[T byte | rune](char T) bool {
	return char >= 'a' && char <= 'z'
}
```

### IsUpper
```go
import ()
```

```go
func IsUpper[T byte | rune](char T) bool {
	return char >= 'A' && char <= 'Z'
}
```
