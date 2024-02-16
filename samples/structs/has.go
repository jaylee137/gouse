package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Has_Person struct {
	Name  string
	Age   int
	Email string
}

func structHas() {
	person := Has_Person{
		Name:  "Example",
		Age:   40,
		Email: "",
	}

	has := structs.Has(person, "Email")
	fmt.Printf("Has: %+v\n", has)

	hasEmpty := structs.HasEmpty(person, "Email")
	fmt.Printf("Has empty: %+v\n", hasEmpty)
}
