package strings

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/strings"
)

func SampleStringSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", strings.Split(str, "l"))
}

func SampleStringWords() {
	var str = "hello world"
	println("Split string to array of words: ", strings.Words(str))

}

func SampleStringReverse() {
	var str = "hello world"
	println("Reverse string: ", strings.Reverse(str))
}

func SampleStringLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", strings.Lowers(str))
	println("Lower string (byte): ", strings.Lower(str[0]))
	println("Lower first string: ", strings.LowerFirst(str))
}

func SampleStringUpper() {
	var str = "hello world"
	println("Upper string (string): ", strings.Uppers(str))
	println("Upper string (byte): ", strings.Upper(str[0]))
	println("Upper first string: ", strings.UpperFirst(str))
}

func SampleStringRepeat() {
	var str = "hello world"
	println("Repeat string: ", strings.Repeat(str, 3))
}

func SampleStringTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", strings.Truncate(str, 5))
	println("Truncate string (custom): ", strings.Truncate(str, 5, "***"))
}

func SampleStringReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", strings.Replace(str, "world", "golang"))
}

func SampleStringTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", strings.Trim(str))
	println("Trim left string: ", strings.LTrim(str))
	println("Trim right string: ", strings.RTrim(str))
}

func SampleStringTrimBlank() {
	println("Trim blank string: ", strings.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", strings.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", strings.TrimBlank("   hello world, this is world   \n"))
}

func SampleStringTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", strings.TrimPrefix(str, "   "))
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}

func SampleStringTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}

func SampleStringJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", strings.Join(str, "-"))
}

func SampleStringConcat() {
	println("Concat string: ", strings.Concat("hello", "world"))
}

func SampleStringSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", strings.SubStr(str, 0, 5))
	println("Sub string: ", strings.SubStr(str, 0, 1))
	println("Sub string (only start): ", strings.SubStr(str, -5))
	println("Sub string (with negative index): ", strings.SubStr(str, -5, -1))
}

func SampleStringSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", strings.Slice(str, 0, 5))
	println("Slice string: ", strings.Slice(str, 0, 1))
	println("Slice string (only start): ", strings.Slice(str, -5))
	println("Slice string (not parameters): ", strings.Slice(str))
	println("Slice string (with negative index): ", strings.Slice(str, -5, -1))
}

func SampleStringSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	println("Splice string (with replace): ", strings.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", strings.Splice(str, 1, 5, "golang1", "golang2"))
}

func SampleStringStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", strings.StartsWith(str, "hello"))
}

func SampleStringEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", strings.EndsWith(str, "world"))
}

func SampleStringEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", strings.Escape(str))
}

func SampleStringUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", strings.Unescape(str))
}

func SampleStringPad() {
	var str = "hello world"
	println("Pad-left string: ", strings.PadStart(str, 20, '$'))
	println("Pad-right string: ", strings.PadEnd(str, 20, '@'))
}
