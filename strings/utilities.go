package strings

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/constants"
	"github.com/thuongtruong1009/gouse/number"
)

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

func FIndex(s string, substr string) int {
	firstIndex, _ := Index(s, substr)
	return firstIndex
}

func LIndex(s string, substr string) int {
	_, lastIndex := Index(s, substr)
	return lastIndex
}

func RandomStr(n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += string(constants.ChainStr[number.Random(0, len(constants.ChainStr)-1)])
	}

	return result
}

func RandomNum(n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += string(constants.ChainNum[number.Random(0, len(constants.ChainNum)-1)])
	}

	return result
}

func At(s string, index int) string {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return ""
	}

	return string(s[index])
}

func CodePointAt(s string, index int) int {
	if index < 0 {
		index = len(s) + index
	}

	if index < 0 || index >= len(s) {
		return -1
	}

	return int(s[index])
}

func CodePoint(input string) []int {
	asciiValues := make([]int, len(input))

	for i, char := range input {
		asciiValues[i] = int(char)
	}

	return asciiValues
}

func FromCodePointAt(codePoint int) string {
	return string(rune(codePoint))
}

func FromCodePoint(codePoint ...int) string {
	var result string

	for _, v := range codePoint {
		result += fmt.Sprintf("%c", v)
	}

	return result
}
