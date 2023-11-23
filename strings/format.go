package strings

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

func IsLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func Reverse(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func LowerS(s string) string {
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

func LowerB(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}

func UpperS(s string) string {
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

func UpperB(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func Repeat(s string, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += s
	}

	return result
}

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

func Replace(s string, old string, new string, count ...int) string {
	if len(count) == 0 {
		count = append(count, -1)
	}

	var result string
	var temp string

	for _, v := range s {
		if string(v) == old {
			temp += string(v)
		} else {
			if len(temp) == len(old) {
				result += new
				temp = ""
			}
			result += temp + string(v)
			temp = ""
		}
	}

	if len(temp) == len(old) {
		result += new
		temp = ""
	}

	return result
}


// Below are testing functions

func PadStart(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}

	return Repeat(pad, length-len(s)) + s
}

func PadEnd(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}

	return s + Repeat(pad, length-len(s))
}

func StartCase(s string) string {
	var result string
	var flag bool

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			flag = true
			result += string(v)
		} else if v >= 'a' && v <= 'z' {
			if flag {
				result += string(v)
				flag = false
			} else {
				result += string(v - 32)
			}
		} else {
			flag = true
			result += " "
		}
	}

	return result
}

func Index(s string, substr string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == substr[0] && len(s)-i >= len(substr) {
			if s[i:i+len(substr)] == substr {
				return i
			}
		}
	}

	return -1
}

func LastIndex(s string, substr string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == substr[0] && len(s)-i >= len(substr) {
			if s[i:i+len(substr)] == substr {
				return i
			}
		}
	}

	return -1
}

func Includes(s string, substr string) bool {
	return Index(s, substr) != -1
}

func StartsWith(s string, substr string) bool {
	return Index(s, substr) == 0
}

func EndsWith(s string, substr string) bool {
	return LastIndex(s, substr) == len(s)-len(substr)
}


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

func TrimPrefix(s string, prefix string) string {
	if len(s) < len(prefix) {
		return s
	}

	if s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}

	return s
}

func TrimSuffix(s string, suffix string) string {
	if len(s) < len(suffix) {
		return s
	}

	if s[len(s)-len(suffix):] == suffix {
		return s[:len(s)-len(suffix)]
	}

	return s
}

func TrimSpace(s string) string {
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