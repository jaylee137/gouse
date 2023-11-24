package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/strings"
)

func strCapitalize() {
	var str = "hello world"
	fmt.Println(strings.Capitalize(str))
}

func strCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	fmt.Println("Convert string to Camel Case: ", strings.CamelCase(str))
}

func strSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	fmt.Println("Convert string to Snake Case: ", strings.SnakeCase(str1), strings.SnakeCase(str2))
}

func strKebabCase() {
	var str = "hello world"
	fmt.Println("Convert string to Kebab Case: ", strings.KebabCase(str))
}

func strIsLetter() {
	var str = "hello world"
	fmt.Println("Check is letter character: ", strings.IsLetter(str[0]))
}

func strIsDigit() {
	var str = "1hello world"
	fmt.Println("Check is number character: ", strings.IsDigit(str[0]))
}

func strIncludes() {
	var str = "hello world, this is world"
	fmt.Println("Check substring in string: ", strings.Includes(str, "world"))
}

func strIsLower() {
	var str = "hELLO WORLD"
	fmt.Println("Check is lower string: ", strings.IsLower(str[0]))
}

func strIsUpper() {
	var str = "Hello world"
	fmt.Println("Check is upper string: ", strings.IsUpper(str[0]))
}

func strSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", strings.Split(str, "l"))
}

func strWords() {
	var str = "hello world"
	fmt.Println("Split string to array of words: ", strings.Words(str))

}

func strReverse() {
	var str = "hello world"
	fmt.Println("Reverse string: ", strings.Reverse(str))
}

func strLower() {
	var str = "HELLO WORLD"
	fmt.Println("Lower string (string): ", strings.Lowers(str))
	fmt.Println("Lower string (byte): ", strings.Lower(str[0]))
	fmt.Println("Lower first string: ", strings.LowerFirst(str))
}

func strUpper() {
	var str = "hello world"
	fmt.Println("Upper string (string): ", strings.Uppers(str))
	fmt.Println("Upper string (byte): ", strings.Upper(str[0]))
	fmt.Println("Upper first string: ", strings.UpperFirst(str))
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

func strReplace() {
	var str = "hello world, this is world"
	fmt.Println("Replace string: ", strings.Replace(str, "world", "golang"))
}

func strTrim() {
	var str = "   hello world, this is world   "
	fmt.Println("Trim string: ", strings.Trim(str))
	fmt.Println("Trim left string: ", strings.LTrim(str))
	fmt.Println("Trim right string: ", strings.RTrim(str))
}

func strJoin() {
	var str = []string{"hello", "world"}
	fmt.Println("Join string: ", strings.Join(str, "-"))
}

func strSlice() {
	var str = "hello world, this is world"
	fmt.Println("Slice string: ", strings.Slice(str, 0, 5))
}

func strSplice() {
	var str = "hello world, this is world"
	fmt.Println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	fmt.Println("Splice string (with replace): ", strings.Splice(str, 0, 5, "golang"))
}

func strStartsWith() {
	var str = "hello world, this is world"
	fmt.Println("Starts with: ", strings.StartsWith(str, "hello"))
}

func strEndsWith() {
	var str = "hello world, this is world"
	fmt.Println("Ends with: ", strings.EndsWith(str, "world"))
}

func strEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	fmt.Println("Escape string: ", strings.Escape(str))
}

func strUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	fmt.Println("Unescape string: ", strings.Unescape(str))
}

func strPad() {
	var str = "hello world"
	// fmt.Println("Pad string: ", strings.Pad(str, 20, '*'))
	fmt.Println("Pad-left string: ", strings.PadStart(str, 20, '$'))
	fmt.Println("Pad-right string: ", strings.PadEnd(str, 20, '@'))
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

func strIndex() {
	var str = "hello world, this is world"
	fmt.Println("First index of substring start at: ", strings.FIndex(str, "world"))
	fmt.Println("Last Index of substring end at: ", strings.LIndex(str, "world"))
}

func strRandom() {
	fmt.Println("Random string: ", strings.Random(10))
}

func StringExample() {
	strCapitalize()
	strCamelCase()
	strSnakeCase()
	strKebabCase()
	
	strIsLetter()
	strIsDigit()
	strIncludes()
	strIsLower()
	strIsUpper()

	strSplit()
	strWords()
	strReverse()
	strLower()
	strUpper()
	strRepeat()
	strTruncate()
	strReplace()
	strTrim()
	strJoin()
	strSlice()
	strSplice()
	strStartsWith()
	strEndsWith()
	strEscape()
	strUnescape()
	strPad()
	
	strCount()
	strLines()
	strIndex()
	strRandom()
}