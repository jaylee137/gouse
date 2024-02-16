package regex

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/regex"
)

// Note: regex pattern is not include one of (^, $, /g)

func regexMatch() {
	fmt.Println("Match string with regex: ", regex.Match(`[A-Z]`, "Hello World 123"))
}
