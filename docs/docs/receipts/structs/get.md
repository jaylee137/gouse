# Get

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/structs"
)
```
## Functions


### SampleStructGet

```go
func SampleStructGet() {
	person := Get_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	name := structs.Get(person, "Name")
	fmt.Printf("Name: %s\n", name)
}
```
