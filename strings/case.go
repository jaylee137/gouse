package strings

func Capitalize(sentence string) string {
	sentenceBytes := []byte(sentence)

	for i := 0; i < len(sentenceBytes); i++ {
		if (i == 0 || sentenceBytes[i-1] == ' ') && IsLetter(sentenceBytes[i]) {
			sentenceBytes[i] = UpperB(sentenceBytes[i])
		}
	}

	return string(sentenceBytes)
}

// func CamelCase(s string) string {
// 	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

//     // Replace all underscores with spaces
//     s = strings.ReplaceAll(s, "_", " ")

//     // Title case s
//     s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

//     // Remove all spaces
//     s = strings.ReplaceAll(s, " ", "")

//     // Lowercase the first letter
//     if len(s) > 0 {
//         s = strings.ToLower(s[:1]) + s[1:]
//     }

//     return s
// }

func SnakeCase(s string) string {
	var result string

	// newStr := s.SplitCase(" ", s)
	return result[1:]
}

func KebabCase(s string) string {
	var result string

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += "-" + string(v+32)
		} else {
			result += string(v)
		}
	}

	return result
}