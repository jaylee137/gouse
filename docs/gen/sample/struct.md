# Struct

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/structs"
)
```
## Functions


### structMerge

```go
func structMerge() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	address := Address{
		City:    "New York",
		Street:  "123 Main St",
		ZipCode: "10001",
	}

	merged := structs.Merge(person, address)

	fmt.Printf("Struct after merged: %+v\n", merged)

	fmt.Println("Name:", merged.(map[string]interface{})["Name"])

	fmt.Println("City:", merged.(map[string]interface{})["City"])
}
```

### structRemove

```go
func structRemove() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Struct after removed field: %+v\n", structs.Remove(person, "Email"))
}
```

### structAdd

```go
func structAdd() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	newFields := map[string]interface{}{
		"Address": "123 Main St",
		"Phone":   "555-1234",
	}
	result := structs.Add(person, newFields)

	fmt.Printf("Struct after adding fields: %+v\n", result)
}
```

### structSet

```go
func structSet() {
	person := &Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	structs.Set(person, "Name", "Updated Name")

	fmt.Printf("Struct after setting field: %+v\n", person)
}
```

### structGet

```go
func structGet() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	name := structs.Get(person, "Name")
	fmt.Printf("Name: %s\n", name)
}
```

### structClone

```go
func structClone() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Original: %+v\n", person)

	clone := structs.Clone(person)

	updateClone := clone.(Person)
	updateClone.Name = "Updated Name"
	fmt.Printf("Clone: %+v\n", updateClone)
}
```

### structHas

```go
func structHas() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "",
	}

	has := structs.Has(person, "Email")
	fmt.Printf("Has: %+v\n", has)

	hasEmpty := structs.HasEmpty(person, "Email")
	fmt.Printf("Has empty: %+v\n", hasEmpty)
}
```

### StructExample

```go
func StructExample() {
	structMerge()
	structRemove()
	structAdd()
	structSet()
	structGet()

	structClone()
	structHas()
}
```
