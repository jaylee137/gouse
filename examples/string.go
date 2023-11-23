package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/strings"
)

func strCapitalize() {
	var str = "hello world"
	fmt.Println(strings.Capitalize(str))
}

// func strCamelCase() {
// 	var str = "hello world"
// 	fmt.Println(strings.CamelCase(str))
// }

// func strSnakeCase() {
// 	var str = "hello world"
// 	fmt.Println(strings.SnakeCase(str))
// }

func strSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", strings.Split(str, " "))
}

func strIsLetter() {
	var str = "hello world"
	fmt.Println("Check is letter character: ", strings.IsLetter(str[0]))
}

func strReverse() {
	var str = "hello world"
	fmt.Println("Reverse string: ", strings.Reverse(str))
}

func strLower() {
	var str = "HELLO WORLD"
	fmt.Println("Lower string (string): ", strings.LowerS(str))
	fmt.Println("Lower string (byte): ", strings.LowerB(str[0]))
}

func strUpper() {
	var str = "hello world"
	fmt.Println("Upper string (string): ", strings.UpperS(str))
	fmt.Println("Upper string (byte): ", strings.UpperB(str[0]))
}

func strRepeat() {
	var str = "hello world"
	fmt.Println("Repeat string: ", strings.Repeat(str, 3))
}

func strTruncate() {
	var str = "hello world"
	fmt.Println("Truncate string (default): ", strings.Truncate(str, 5))
	fmt.Println("Truncate string (custom): ", strings.Truncate(str, 5, "***"))
}

func strCount() {
	var str = "hello world wo wo"
	fmt.Println("Count words/substr in string (default): ", strings.Count(str))
	fmt.Println("Count words/substr in string (with char): ", strings.Count(str, "wo"))
}

func strLines() {
	var str = "hello world\nwo wo"
	fmt.Println("Count lines of string: ", strings.Lines(str))
}

func strReplace() {
	var str = "hello world"
	fmt.Println("Replace string: ", strings.Replace(str, "world", "golang"))
	fmt.Println("Replace string: ", strings.Replace(str, "world", "golang", 1))
}

func StringExample() {
	// strCapitalize()
	// strCamelCase()
	// strSnakeCase()

	// strSplit()
	// strIsLetter()
	// strReverse()
	// strLower()
	// strUpper()
	// strRepeat()
	// strTruncate()
	// strCount()
	// strLines()
	strReplace()
}