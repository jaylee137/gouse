
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

### minmax
```go
import ()
```

```go
func minmax[T comparable](arr []T, less func(T, T) bool) T {
	if len(arr) == 0 {
		panic("Empty array")
	}

	max := arr[0]
	for _, item := range arr {
		if less(max, item) {
			max = item
		}
	}
	return max
}
```

### Min
```go
import ()
```

```go
func Min[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return minmax[T](arr, func(a, b T) bool {
		return a > b
	})
}
```

### Max
```go
import ()
```

```go
func Max[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](arr []T) T {
	return minmax[T](arr, func(a, b T) bool {
		return a < b
	})
}
```

### TestMin
```go
import (
	"testing"
)
```

```go
func TestMin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	minExpected := 1
	if Min(arr) != minExpected {
		t.Errorf("Expected %d, got %d", minExpected, Min(arr))
	}

	maxExpected := 5
	if Max(arr) != maxExpected {
		t.Errorf("Expected %d, got %d", maxExpected, Max(arr))
	}
}
```

### Includes
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Includes[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
```

### Equal
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Equal[T comparable](a, b T) bool {
	return a == b
}
```

### Sum
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Sum[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}
```

### Most
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Most[T comparable](arr []T) T {
	var most = make(map[T]int)
	max := arr[0]
	for _, v := range arr {
		most[v] = most[v] + 1
		if most[max] < most[v] {
			max = v
		}
	}
	return max
}
```

### Chunk
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Chunk[T comparable](array []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}
```

### Diff
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Diff[T comparable](a, b []T) []T {
	var diff []T
	for _, v := range a {
		if !Includes(b, v) {
			diff = append(diff, v)
		}
	}
	return diff
}
```

### Drop
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Drop[T comparable](arr []T, n ...int) []T {
	if len(n) == 0 {
		n = append(n, 1)
	}

	return arr[n[0]:]
}
```

### Index
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Index[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
```

### Merge
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Merge[T comparable](arr ...[]T) []T {
	var merged []T
	for _, v := range arr {
		merged = append(merged, v...)
	}
	return merged
}
```

### Compact
```go
import (
	"github.com/thuongtruong1009/gouse/types"
)
```

```go
func Compact[T interface{}](arr []T) []T {
	var compact []T
	for _, v := range arr {
		if !types.IsZero(v) && !types.IsNil(v) && !types.IsEmpty(v) && !types.IsUndefined(v) && !types.IsBool(v) {
			compact = append(compact, v)
		}
	}
	return compact
}
```
