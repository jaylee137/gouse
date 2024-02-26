# Clone

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/structs"
)
```
## Functions


### SampleStructClone

```go
func SampleStructClone() {
	person := Clone_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Original: %+v\n", person)

	clone := structs.Clone(person)

	updateClone := clone.(Clone_Person)
	updateClone.Name = "Updated Name"
	fmt.Printf("Clone: %+v\n", updateClone)
}
```
