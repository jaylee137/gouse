package regex

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/regex"
)

// Note: regex pattern is not include one of (^, $, /g)

func SampleRegexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := regex.MatchIndex(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
