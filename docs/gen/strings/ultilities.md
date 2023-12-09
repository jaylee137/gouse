# Ultilities


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

### Split
```go
import (
	"html"
)
```

```go
func Split(s string, separator string) []string {
	var result []string
	var temp string

	for _, v := range s {
		if string(v) == separator {
			result = append(result, temp)
			temp = ""
		} else {
			temp += string(v)
		}
	}

	result = append(result, temp)

	return result
}
```

### Words
```go
import (
	"html"
)
```

```go
func Words(s string) []string {
	var result []string

	for _, v := range s {
		if v == ' ' || v == '\n' || v == '\t' {
			continue
		} else {
			result = append(result, string(v))
		}
	}

	return result
}
```

### Reverse
```go
import (
	"html"
)
```

```go
func Reverse(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
	}

	return result
}
```

### Lowers
```go
import (
	"html"
)
```

```go
func Lowers(s string) string {
	var result string

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}

	return result
}
```

### Lower
```go
import (
	"html"
)
```

```go
func Lower[T byte | rune](char T) T {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}
```

### LowerFirst
```go
import (
	"html"
)
```

```go
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Lower(s[0])) + s[1:]
}
```

### Uppers
```go
import (
	"html"
)
```

```go
func Uppers(s string) string {
	var result string

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			result += string(v - 32)
		} else {
			result += string(v)
		}
	}

	return result
}
```

### Upper
```go
import (
	"html"
)
```

```go
func Upper[T byte | rune](char T) T {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}
```

### UpperFirst
```go
import (
	"html"
)
```

```go
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Upper(s[0])) + s[1:]
}
```

### Repeat
```go
import (
	"html"
)
```

```go
func Repeat(s string, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += s
	}

	return result

}
```

### Truncate
```go
import (
	"html"
)
```

```go
func Truncate(s string, length int, omission ...string) string {
	if len(s) <= length {
		return s
	}

	if len(omission) == 0 {
		return s[:length] + "..."
	}

	if len(omission) > length {
		return s
	}

	return s[:length] + omission[0]
}
```

### Replace
```go
import (
	"html"
)
```

```go
func Replace(s string, old string, new string) string {
	var result string

	for i := 0; i < len(s); {
		// Check if the remaining part of the input matches the old
		if len(s)-i >= len(old) && s[i:i+len(old)] == old {
			// If there's a match, append the newSubstr to the result and move the index past the old
			result += new
			i += len(old)
		} else {
			// If no match, append the current character to the result and move to the next character
			result += string(s[i])
			i++
		}
	}

	return result
}
```

### Trim
```go
import (
	"html"
)
```

```go
func Trim(s string) string {
	var result string
	var flag bool

	for _, v := range s {
		if v != ' ' {
			flag = true
			result += string(v)
		} else if flag {
			result += string(v)
		}
	}

	return result
}
```

### LTrim
```go
import (
	"html"
)
```

```go
func LTrim(s string) string {
	var result string
	var flag bool

	for _, v := range s {
		if v != ' ' {
			flag = true
			result += string(v)
		} else if flag {
			result += string(v)
		}
	}

	return result
}
```

### RTrim
```go
import (
	"html"
)
```

```go
func RTrim(s string) string {
	var result string
	var flag bool

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
			result = string(s[i]) + result
		} else if flag {
			result = string(s[i]) + result
		}
	}

	return result
}
```

### TrimBlank
```go
import (
	"html"
)
```

```go
func TrimBlank(s string) string {
	var result string
	var flag bool

	for _, v := range s {
		if v != ' ' && v != '\n' && v != '\t' {
			flag = true
			result += string(v)
		} else if flag {
			result += string(v)
		}
	}

	return result
}
```

### TrimPrefix
```go
import (
	"html"
)
```

```go
func TrimPrefix(s string, prefix string) string {
	if len(s) < len(prefix) {
		return s
	}

	if s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}

	return s
}
```

### TrimSuffix
```go
import (
	"html"
)
```

```go
func TrimSuffix(s string, suffix string) string {
	if len(s) < len(suffix) {
		return s
	}

	if s[len(s)-len(suffix):] == suffix {
		return s[:len(s)-len(suffix)]
	}

	return s
}
```

### Join
```go
import (
	"html"
)
```

```go
func Join(s []string, separator string) string {
	var result string

	for i, v := range s {
		result += v
		if i != len(s)-1 {
			result += separator
		}
	}

	return result
}
```

### Concat
```go
import (
	"html"
)
```

```go
func Concat(s ...string) string {
	var result string

	for _, v := range s {
		result += v
	}

	return result
}
```

### Slice
```go
import (
	"html"
)
```

```go
func Slice(s string, start int, end int) string {
	if start < 0 {
		start = len(s) + start
	}

	if end < 0 {
		end = len(s) + end
	}

	if end > len(s) {
		end = len(s)
	}

	if start > end {
		return ""
	}

	return s[start:end]
}
```

### Splice
```go
import (
	"html"
)
```

```go
func Splice(s string, start int, deleteCount int, items ...string) string {
	if start < 0 {
		start = len(s) + start
	}

	if deleteCount < 0 {
		deleteCount = 0
	}

	if start > len(s) {
		start = len(s)
	}

	if deleteCount > len(s)-start {
		deleteCount = len(s) - start
	}

	return s[:start] + Join(items, "") + s[start+deleteCount:]
}
```

### Escape
```go
import (
	"html"
)
```

```go
func Escape(s string) string {
	var result string

	for _, v := range s {
		if v == '\\' || v == '/' || v == '[' || v == ']' || v == '{' || v == '}' || v == '(' || v == ')' || v == '*' || v == '+' || v == '?' || v == '.' || v == '^' || v == '$' || v == '|' {
			result += "\\" + string(v)
		}

		switch v {
		case '&':
			result += "&amp;"
		case '<':
			result += "&lt;"
		case '>':
			result += "&gt;"
		case '"':
			result += "&quot;"
		case '\'':
			result += "&#39;"
		case '\u2013': // En dash
			result += "&ndash;"
		case '\u2014': // Em dash
			result += "&mdash;"
		case '\u2018', '\u2019': // Left and right single quotation mark
			result += "&lsquo;"
		case '\u201C', '\u201D': // Left and right double quotation mark
			result += "&ldquo;"
		case '\u00A9': // Copyright symbol
			result += "&copy;"
		case '\u00AE': // Registered trademark symbol
			result += "&reg;"
		default:
			result += string(v)
		}
	}

	return result
}
```

### Unescape
```go
import (
	"html"
)
```

```go
func Unescape(s string) string {
	s = html.UnescapeString(s)

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if i+1 < len(s) {
				i++
			}
		}
	}

	return s
}
```

### pad
```go
import (
	"html"
)
```

```go
func pad(s string, length int, padChar byte) string {
	if len(s) >= length {
		return s
	}

	padding := make([]byte, length-len(s))
	for i := range padding {
		padding[i] = padChar
	}

	return string(padding)
}
```

### PadStart
```go
import (
	"html"
)
```

```go
func PadStart(s string, length int, padChar byte) string {
	return pad(s, length, padChar) + s
}
```

### PadEnd
```go
import (
	"html"
)
```

```go
func PadEnd(s string, length int, padChar byte) string {
	return s + pad(s, length, padChar)
}
```

### Count
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func Count(s string, substr ...string) int {
	if len(substr) == 0 {
		return len(s)
	}

	var count int

	for i := 0; i < len(s); i++ {
		if s[i] == substr[0][0] && len(s)-i >= len(substr[0]) {
			if s[i:i+len(substr[0])] == substr[0] {
				count++
				i += len(substr[0]) - 1
			}
		}
	}

	return count
}
```

### Lines
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func Lines(s string) int {
	var result []string
	var temp string

	for _, v := range s {
		if v == '\n' {
			result = append(result, temp)
			temp = ""
		} else {
			temp += string(v)
		}
	}

	result = append(result, temp)

	return len(result)
}
```

### FIndex
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func FIndex(s string, substr string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == substr[0] && len(s)-i >= len(substr) {
			if s[i:i+len(substr)] == substr {
				return i
			}
		}
	}

	return -1
}
```

### LIndex
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func LIndex(s string, substr string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == substr[0] && len(s)-i >= len(substr) {
			if s[i:i+len(substr)] == substr {
				return i
			}
		}
	}

	return -1
}
```

### Random
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func Random(n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += string(constants.KeyStr[number.Random(0, len(constants.KeyStr)-1)])
	}

	return result
}
```
