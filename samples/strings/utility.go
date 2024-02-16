package strings

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/strings"
)

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
	println("Random chain string: ", strings.RandomStr(10))

	println("Random chain number: ", strings.RandomNum(6))
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
