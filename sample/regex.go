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

func regexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := regex.MatchIndex(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
