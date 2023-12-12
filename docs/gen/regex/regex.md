# Regex


### IsMatch
```go
import (
	"regexp"
)
```

```go
func IsMatch(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}
```

### Match
```go
import (
	"regexp"
)
```

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
