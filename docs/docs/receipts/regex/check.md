# Check

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/regex"
)
```
## Functions


### SampleRegexIsMatch

```go
func SampleRegexIsMatch() {
	fmt.Println("Match string with regex: ", regex.IsMatch(`[a-z]+\s[a-z]+`, "hello world"))
}
```
