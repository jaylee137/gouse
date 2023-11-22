package strings

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func toUpper(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func Capitalize(sentence string) string {
	sentenceBytes := []byte(sentence)

	for i := 0; i < len(sentenceBytes); i++ {
		if (i == 0 || sentenceBytes[i-1] == ' ') && isLetter(sentenceBytes[i]) {
			sentenceBytes[i] = toUpper(sentenceBytes[i])
		}
	}

	return string(sentenceBytes)
}

func CamelCase(s string) string {
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
		}
	}

	return result
}