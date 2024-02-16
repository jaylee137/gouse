package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Clone_Person struct {
	Name  string
	Age   int
	Email string
}

func structClone() {
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
