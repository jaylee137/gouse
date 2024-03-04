# Check

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/strings"
)
```
## Functions


### SampleStringIsLetter

```go
func SampleStringIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", strings.IsLetter(str[0]))
}
```

### SampleStringIsDigit

```go
func SampleStringIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", strings.IsDigit(str[0]))
}
```

### SampleStringIncludes

```go
func SampleStringIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", strings.Includes(str, "world"))
}
```

### SampleStringIsLower

```go
func SampleStringIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", strings.IsLower(str[0]))
}
```

### SampleStringIsUpper

```go
func SampleStringIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", strings.IsUpper(str[0]))
}
```
