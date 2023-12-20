# Regex

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/regex"
)
```
## Functions


### regexIsMatch

```go
func regexIsMatch() {
	fmt.Println("Match string with regex: ", regex.IsMatch(`[a-z]+\s[a-z]+`, "hello world"))
}
```

### regexMatch

```go
func regexMatch() {
	fmt.Println("Match string with regex: ", regex.Match(`[A-Z]`, "Hello World 123"))
}
```

### regexMatchIndex

```go
func regexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := regex.MatchIndex(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
```

### RegexExample

```go
func RegexExample() {
	regexIsMatch()
	regexMatch()
	regexMatchIndex()
}
```
