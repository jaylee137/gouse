# Case

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/strings"
)
```
## Functions


### SampleStringCapitalize

```go
func SampleStringCapitalize() {
	var str = "hello world"
	println(strings.Capitalize(str))
}
```

### SampleStringCamelCase

```go
func SampleStringCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", strings.CamelCase(str))
}
```

### SampleStringSnakeCase

```go
func SampleStringSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", strings.SnakeCase(str1), strings.SnakeCase(str2))
}
```

### SampleStringKebabCase

```go
func SampleStringKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", strings.KebabCase(str))
}
```
