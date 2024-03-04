package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Get_Person struct {
	Name  string
	Age   int
	Email string
}

func SampleStructGet() {
	person := Get_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	name := structs.Get(person, "Name")
	fmt.Printf("Name: %s\n", name)
}
