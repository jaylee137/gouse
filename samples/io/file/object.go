package file

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thuongtruong1009/gouse/io"
)

func ioFileObj() {
	type User struct {
		Name string
		Age  int
	}

	exampleFile := "data.json"

	// read file
	// data, err := io.ReadFileObj[User](exampleFile)
	// if err != nil {
	// 	println(err.Error())
	// }
	// fmt.Printf("data: %+v\n", data)

	// // write file
	// updateData := append(data, User{
	// 	Name: fmt.Sprintf("name %d", len(data)+1),
	// 	Age:  len(data) + 1,
	// })

	data, err := os.ReadFile(exampleFile)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("data: %+v\n", data)

	// unmarshal data into a slice of User
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		println(err.Error())
		return
	}

	// create a new user
	newUser := User{
		Name: fmt.Sprintf("name %d", len(users)+1),
		Age:  len(users) + 1,
	}

	// append the new user to the slice
	users = append(users, newUser)

	// marshal the updated slice back to JSON
	updatedData, err := json.Marshal(users)
	if err != nil {
		println(err.Error())
		return
	}

	// write the updated data back to the file
	// if err := ioutil.WriteFile(exampleFile, updatedData, 0644); err != nil {
	// 	println(err.Error())
	// 	return
	// }

	err2 := io.WriteFileObj(exampleFile, updatedData)
	if err2 != nil {
		println(err2.Error())
	}
	println("data written")
}
