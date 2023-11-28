package strings

import "html"

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

func Reverse(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
	}

	return result
}

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

func Lower[T byte | rune](char T) T {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}

func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Lower(s[0])) + s[1:]
}

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

func Upper[T byte | rune](char T) T {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return string(Upper(s[0])) + s[1:]
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

func Concat(s ...string) string {
	var result string

	for _, v := range s {
		result += v
	}

	return result
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

// func Pad(s string, length int, pad ...string) string {
// 	if len(s) >= length {
// 		return s
// 	}

// 	if len(pad) == 0 {
// 		pad = append(pad, " ")
// 	}

// 	return Repeat(pad, (length-len(s))/2) + s + Repeat(pad, (length-len(s)+1)/2)
// }

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

func PadStart(s string, length int, padChar byte) string {
	return pad(s, length, padChar) + s
}

func PadEnd(s string, length int, padChar byte) string {
	return s + pad(s, length, padChar)
}
