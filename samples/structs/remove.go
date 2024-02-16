package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Remove_Person struct {
	Name  string
	Age   int
	Email string
}

func structRemove() {
	person := Remove_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Struct after removed field: %+v\n", structs.Remove(person, "Email"))
}
