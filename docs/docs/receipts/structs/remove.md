# Remove

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/structs"
)
```
## Functions


### SampleStructRemove

```go
func SampleStructRemove() {
	person := Remove_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Struct after removed field: %+v\n", structs.Remove(person, "Email"))
}
```
