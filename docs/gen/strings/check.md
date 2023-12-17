# Check


### Capitalize

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
func SnakeCase(s string) string {
	return concatCase(s, "_")
}
```

### KebabCase

```go
func KebabCase(s string) string {
	return concatCase(s, "-")
}
```

### IsLetter

```go
func IsLetter[T byte | rune](char T) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
```

### IsDigit

```go
func IsDigit[T byte | rune](char T) bool {
	return char >= '0' && char <= '9'
}
```

### Includes

```go
func Includes(s string, substr string) bool {
	return FIndex(s, substr) != -1
}
```

### StartsWith

```go
func StartsWith(s string, substr string) bool {
	return FIndex(s, substr) == 0
}
```

### EndsWith

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
func IsLower[T byte | rune](char T) bool {
	return char >= 'a' && char <= 'z'
}
```

### IsUpper

```go
func IsUpper[T byte | rune](char T) bool {
	return char >= 'A' && char <= 'Z'
}
```
