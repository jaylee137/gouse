package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/strings"
)

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
	fmt.Println("Split string by separator: ", strings.Split(str, "l"))
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

func strTrimBlank() {
	println("Trim blank string: ", strings.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", strings.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", strings.TrimBlank("   hello world, this is world   \n"))
}

func strTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", strings.TrimPrefix(str, "   "))
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}

func strTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}

func strJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", strings.Join(str, "-"))
}

func strConcat() {
	println("Concat string: ", strings.Concat("hello", "world"))
}

func strSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", strings.SubStr(str, 0, 5))
	println("Sub string: ", strings.SubStr(str, 0, 1))
	println("Sub string (only start): ", strings.SubStr(str, -5))
	println("Sub string (with negative index): ", strings.SubStr(str, -5, -1))
}

func strSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", strings.Slice(str, 0, 5))
	println("Slice string: ", strings.Slice(str, 0, 1))
	println("Slice string (only start): ", strings.Slice(str, -5))
	println("Slice string (not parameters): ", strings.Slice(str))
	println("Slice string (with negative index): ", strings.Slice(str, -5, -1))
}

func strSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	println("Splice string (with replace): ", strings.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", strings.Splice(str, 1, 5, "golang1", "golang2"))
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

	f1, l1 := strings.Index(str, "l")
	fmt.Printf("First index start at: %d, end at: %d\n", f1, l1)

	f2, l2 := strings.Index(str, "world")
	fmt.Printf("First index start at: %d, end at: %d\n", f2, l2)

	f3 := strings.FIndex(str, "l")
	fmt.Println("First index start at: ", f3)

	l3 := strings.LIndex(str, "world")
	fmt.Println("Last index start at: ", l3)

	f4, l4 := strings.Index(str, "oo")
	if f4 == -1 || l4 == -1 {
		fmt.Println("Not found")
	}

	if strings.FIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}

	if strings.LIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}
}

func strRandom() {
	println("Random string: ", strings.Random(10))
}

func strAt() {
	var str = "hello world"
	println("At string: ", strings.At(str, 1))
	println("At string: ", strings.At(str, -5))
}

func strCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", strings.CodePointAt(str, 1))
	println("Code point at string: ", strings.CodePointAt(str, -5))
}

func strCodePoint() {
	asciiValues := strings.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}

func strFromCodePointAt() {
	println("From code point at string: ", strings.FromCodePointAt(9733))
	println("From code point at string: ", strings.FromCodePointAt(9731))
}

func strFromCodePoint() {
	println("From code point string: ", strings.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
