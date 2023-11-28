package strings

func IsLetter[T byte | rune](char T) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func IsDigit[T byte | rune](char T) bool {
	return char >= '0' && char <= '9'
}

func Includes(s string, substr string) bool {
	return FIndex(s, substr) != -1
}

func StartsWith(s string, substr string) bool {
	return FIndex(s, substr) == 0
}

func EndsWith(s string, substr string) bool {
	return LIndex(s, substr) == len(s)-len(substr)
}

func IsLower[T byte | rune](char T) bool {
	return char >= 'a' && char <= 'z'
}

func IsUpper[T byte | rune](char T) bool {
	return char >= 'A' && char <= 'Z'
}
