package regex

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/regex"
)

// Note: regex pattern is not include one of (^, $, /g)

func regexIsMatch() {
	fmt.Println("Match string with regex: ", regex.IsMatch(`[a-z]+\s[a-z]+`, "hello world"))
}
