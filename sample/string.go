package sample

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

func strSplit() {
	var str = "hello world"
	println("Split string by separator: ", strings.Split(str, "l"))
}

func strWords() {
	var str = "hello world"
	println("Split string to array of words: ", strings.Words(str))

}

func strReverse() {
	var str = "hello world"
	println("Reverse string: ", strings.Reverse(str))
}

func strLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", strings.Lowers(str))
	println("Lower string (byte): ", strings.Lower(str[0]))
	println("Lower first string: ", strings.LowerFirst(str))
}

func strUpper() {
	var str = "hello world"
	println("Upper string (string): ", strings.Uppers(str))
	println("Upper string (byte): ", strings.Upper(str[0]))
	println("Upper first string: ", strings.UpperFirst(str))
}

func strRepeat() {
	var str = "hello world"
	println("Repeat string: ", strings.Repeat(str, 3))
}

func strTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", strings.Truncate(str, 5))
	println("Truncate string (custom): ", strings.Truncate(str, 5, "***"))
}

func strReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", strings.Replace(str, "world", "golang"))
}

func strTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", strings.Trim(str))
	println("Trim left string: ", strings.LTrim(str))
	println("Trim right string: ", strings.RTrim(str))
}

func strJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", strings.Join(str, "-"))
}

func strConcat() {
	println("Concat string: ", strings.Concat("hello", "world"))
}

func strSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", strings.Slice(str, 0, 5))
}

func strSplice() {
	var str = "hello world, this is world"
	println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	println("Splice string (with replace): ", strings.Splice(str, 0, 5, "golang"))
}

func strStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", strings.StartsWith(str, "hello"))
}

func strEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", strings.EndsWith(str, "world"))
}

func strEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", strings.Escape(str))
}

func strUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", strings.Unescape(str))
}

func strPad() {
	var str = "hello world"
	// println("Pad string: ", strings.Pad(str, 20, '*'))
	println("Pad-left string: ", strings.PadStart(str, 20, '$'))
	println("Pad-right string: ", strings.PadEnd(str, 20, '@'))
}

func strCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", strings.Count(str))
	println("Count words/substr in string (with char): ", strings.Count(str, "wo"))
}

func strLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", strings.Lines(str))
}

func strIndex() {
	var str = "hello world, this is world"
	println("First index of substring start at: ", strings.FIndex(str, "world"))
	println("Last Index of substring end at: ", strings.LIndex(str, "world"))
}

func strRandom() {
	println("Random string: ", strings.Random(10))
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
	strConcat()
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
