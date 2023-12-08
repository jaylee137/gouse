
### IndexBy
```go
import ()
```

```go
func IndexBy[T comparable](arr []T, f func(T) bool) int {
	for i, v := range arr {
		if f(v) {
			return i
		}
	}
	return -1
}
```

### KeyBy
```go
import ()
```

```go
func KeyBy[T comparable](arr []T, f func(T) bool) int {
	return IndexBy(arr, f)
}
```

### FilterBy
```go
import ()
```

```go
func FilterBy[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
```

### RejectBy
```go
import ()
```

```go
func RejectBy[T comparable](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if !f(v) {
			res = append(res, v)
		}
	}
	return res
}
```

### FindBy
```go
import ()
```

```go
func FindBy[T comparable](arr []T, f func(T) bool) T {
	for _, v := range arr {
		if f(v) {
			return v
		}
	}
	return arr[0]
}
```

### ForBy
```go
import ()
```

```go
func ForBy[T comparable](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}
```

### MapBy
```go
import ()
```

```go
func MapBy[T comparable, R comparable](arr []T, f func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}
```

### intersectSlice
```go
import ()
```

```go
func intersectSlice[T comparable](a, b []T) []T {
	var intersect []T

	for _, v := range a {
		if Includes(b, v) {
			intersect = append(intersect, v)
		}
	}

	return intersect
}
```

### Intersect
```go
import ()
```

```go
func Intersect[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	intersect := slices[0]

	for _, slice := range slices[1:] {
		intersect = intersectSlice(intersect, slice)
	}

	return intersect
}
```
