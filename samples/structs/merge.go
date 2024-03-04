package structs

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/structs"
)

type Merge_Person struct {
	Name  string
	Age   int
	Email string
}

type Merge_Address struct {
	City    string
	Street  string
	ZipCode string
}

func SampleStructMerge() {
	person := Merge_Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	address := Merge_Address{
		City:    "New York",
		Street:  "123 Main St",
		ZipCode: "10001",
	}

	merged := structs.Merge(person, address)

	fmt.Printf("Struct after merged: %+v\n", merged)

	fmt.Println("Name:", merged.(map[string]interface{})["Name"])

	fmt.Println("City:", merged.(map[string]interface{})["City"])
}
