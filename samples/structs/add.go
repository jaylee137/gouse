package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Add_Person struct {
	Name  string
	Age   int
	Email string
}

func SampleStructAdd() {
	person := Add_Person{
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
