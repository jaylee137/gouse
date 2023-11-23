package strings

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