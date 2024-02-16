package strings

import "github.com/thuongtruong1009/gouse/strings"

func strCapitalize() {
	var str = "hello world"
	println(strings.Capitalize(str))
}

func strCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", strings.CamelCase(str))
}

func strSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", strings.SnakeCase(str1), strings.SnakeCase(str2))
}

func strKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", strings.KebabCase(str))
}
