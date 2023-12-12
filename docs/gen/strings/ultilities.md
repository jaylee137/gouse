# Ultilities


### Capitalize
```go
import ()
```

```go
func Capitalize(sentence string) string {
	sentenceBytes := []byte(sentence)

	for i := 0; i < len(sentenceBytes); i++ {
		if (i == 0 || sentenceBytes[i-1] == ' ') && IsLetter(sentenceBytes[i]) {
			sentenceBytes[i] = Upper(sentenceBytes[i])
		}
	}

	return string(sentenceBytes)
}
```

### CamelCase
```go
import ()
```

```go
func CamelCase(s string) string {
	var result string
	upperNext := false

	for _, char := range s {
		if IsLetter(char) || IsDigit(char) {
			if upperNext {
				result += string(Upper(char))
				upperNext = false
			} else {
				result += string(Lower(char))
			}
		} else {
			// If the character is not a letter or digit, set the flag to capitalize the next valid character.
			upperNext = true
		}
	}

	return result
}
```

### concatCase
```go
import ()
```

```go
func concatCase(s string, sep string) string {
	spliStr := Split(s, " ")
	var result string

	if len(spliStr) == 1 {
		for i, char := range s {
			if i > 0 && (IsUpper(char) || IsDigit(char)) && !IsDigit(s[i-1]) {
				result += sep
			} else if i > 0 && !IsLetter(char) && !IsDigit(char) && string(char) != sep {
				result += sep
				continue
			}
			result += string(Lower(char))
		}

		return result
	} else {
		for i, str := range spliStr {
			if i > 0 {
				result += sep
			}
			result += str
		}

		return result
	}
}
```

### SnakeCase
```go
import ()
```

```go
func SnakeCase(s string) string {
	return concatCase(s, "_")
}
```

### KebabCase
```go
import ()
```

```go
func KebabCase(s string) string {
	return concatCase(s, "-")
}
```

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
	if len(s) < len(substr) {
		return false
	}
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

	if len(separator) == 0 {
		for _, v := range s {
			result = append(result, string(v))
		}
	} else {
		for _, v := range s {
			if string(v) == separator {
				result = append(result, temp)
				temp = ""
			} else {
				temp += string(v)
			}
		}
		result = append(result, temp)
	}

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

	// it is capital case
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
	if len(old) == 0 {
		return s
	}

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

### Trim
```go
import (
	"html"
)
```

```go
func Trim(s string) string {
	tmp := LTrim(s)
	return RTrim(tmp)
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
	start := 0
	end := len(s) - 1

	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}

	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
		end--
	}

	return s[start : end+1]
}
```

### TrimPrefix
```go
import (
	"html"
)
```

```go
func TrimPrefix(s, prefix string) string {
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
func TrimSuffix(s, suffix string) string {
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
func Slice(s string, index ...int) string {
	var start, end int
	if len(index) == 0 {
		return s
	}

	if len(index) == 1 {
		start = index[0]
		end = len(s)
	} else {
		start = index[0]
		end = index[1]
	}

	if start < 0 {
		start = len(s) + start
	}

	if end < 0 {
		end = len(s) + end
	}

	if end > len(s) {
		end = len(s)
	}

	if start > end || start == end {
		return ""
	}

	if start == end-1 {
		return string(s[start])
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
func Splice(s string, start, replaceCount int, items ...string) string {
	if len(items) == 0 {
		return s[:start] + s[start+replaceCount:]
	}

	if start < 0 {
		start = len(s) + start
	}

	if replaceCount < 0 {
		replaceCount = 0
	}

	if start > len(s) {
		start = len(s)
	}

	if replaceCount > len(s)-start {
		replaceCount = len(s) - start
	}

	if replaceCount > len(items) {
		replaceCount = len(items) - 1
	}

	return s[:start] + Join(items, "") + s[start+replaceCount:]
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
func pad(s string, addAmount int, padChar byte) string {
	padding := make([]byte, addAmount-len(s))
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
func PadStart(s string, addAmount int, padChar byte) string {
	if len(s) >= addAmount {
		return s
	}
	return pad(s, addAmount, padChar) + s
}
```

### PadEnd
```go
import (
	"html"
)
```

```go
func PadEnd(s string, addAmount int, padChar byte) string {
	if len(s) >= addAmount {
		return s
	}

	return s + pad(s, addAmount, padChar)
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

### Index
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func Index(s, substr string) (int, int) {
	firstIndex := -1
	lastIndex := -1

	for i := 0; i < len(s); i++ {
		if len(s)-i < len(substr) {
			break
		}

		if s[i:i+len(substr)] == substr {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	return firstIndex, lastIndex
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
	firstIndex, _ := Index(s, substr)
	return firstIndex
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
	_, lastIndex := Index(s, substr)
	return lastIndex
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

### At
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func At(s string, index int) string {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return ""
	}

	return string(s[index])
}
```

### CodePointAt
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func CodePointAt(s string, index int) int {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return -1
	}

	return int(s[index])
}
```

### CodePoint
```go
import (
	"github.com/thuongtruong1009/gouse/constants"

	"github.com/thuongtruong1009/gouse/number"
)
```

```go
func CodePoint(input string) []int {
    asciiValues := make([]int, len(input))

    for i, char := range input {
        asciiValues[i] = int(char)
    }

    return asciiValues
}
```
