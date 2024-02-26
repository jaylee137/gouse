package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Set_Person struct {
	Name  string
	Age   int
	Email string
}

func SampleStructSet() {
	person := &Set_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	structs.Set(person, "Name", "Updated Name")

	fmt.Printf("Struct after setting field: %+v\n", person)
}
