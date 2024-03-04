# Set

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/structs"
)
```
## Functions


### SampleStructSet

```go
func SampleStructSet() {
	person := &Set_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	structs.Set(person, "Name", "Updated Name")

	fmt.Printf("Struct after setting field: %+v\n", person)
}
```
