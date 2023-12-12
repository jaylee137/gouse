package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/regex"
)

// Note: regex pattern is not include one of (^, $, /g)

func regexIsMatch() {
	fmt.Println("Match string with regex: ", regex.IsMatch(`[a-z]+\s[a-z]+`, "hello world"))
}

func regexMatch() {
	fmt.Println("Match string with regex: ", regex.Match(`[A-Z]`, "Hello World 123"))
}

func RegexExample() {
	regexIsMatch()
	regexMatch()
}
