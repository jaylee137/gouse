# Regex

## Imports

```go
import (
	"regexp"
)
```
## Functions


### IsMatch

```go
func IsMatch(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}
```

### Match

```go
func Match(regex, chain string) []string {
	var result []string
	for _, v := range chain {
		if IsMatch(regex, string(v)) {
			result = append(result, string(v))
		}
	}
	return result
}
```

### MatchIndex

```go
func MatchIndex(regex, chain string) int {
	for i, v := range chain {
		if IsMatch(regex, string(v)) {
			return i
		}
	}

	return -1
}
```
