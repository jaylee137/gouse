# Callback


### IndexBy

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
func KeyBy[T comparable](arr []T, f func(T) bool) int {
	return IndexBy(arr, f)
}
```

### FilterBy

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
func ForBy[T comparable](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}
```

### MapBy

```go
func MapBy[T comparable, R comparable](arr []T, f func(T) R) []R {
	var res []R
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}
```
