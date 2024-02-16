package strings

import "github.com/thuongtruong1009/gouse/strings"

func strIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", strings.IsLetter(str[0]))
}

func strIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", strings.IsDigit(str[0]))
}

func strIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", strings.Includes(str, "world"))
}

func strIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", strings.IsLower(str[0]))
}

func strIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", strings.IsUpper(str[0]))
}
