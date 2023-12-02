package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/shared"
	"github.com/thuongtruong1009/gouse/io"
)

func ioClearConsole() {
	io.ClearConsole()
	println("console cleared")
}

func ioOutputColor() {
	io.OutputColor(shared.DEFAULT, "this is default")
	io.OutputColor(shared.WHITE, "this is white")
	io.OutputColor(shared.RED, "this is red")
	io.OutputColor(shared.GREEN, "this is green")
	io.OutputColor(shared.YELLOW, "this is yellow")
	io.OutputColor(shared.BLUE, "this is blue")
	io.OutputColor(shared.MAGENTA, "this is magenta")
	io.OutputColor(shared.CYAN, "this is cyan")
}

func ioFile() {
	type User struct {
		Name string
		Age  int
	}

	exampleData := []User{
		{
			Name: "name 1",
			Age:  1,
		},
		{
			Name: "name 2",
			Age:  2,
		},
	}

	exampleFile := "data.json"

	// is exist file
	err4 := io.IsExistFile(exampleFile)
	if err4 != nil {
		println(err4.Error())
		return
	}
	println("file exist")

	// read file
	data, err := io.ReadFile[*User](exampleFile)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("data: %+v\n", data)

	// write file
	err2 := io.WriteFile(exampleFile, exampleData)
	if err2 != nil {
		println(err2.Error())
		return
	}
	println("data written")
	
	// remove file
	err6 := io.RemoveFile(exampleFile)
	if err6 != nil {
		println(err6.Error())
		return
	}
	println("file removed")

	// is exist dir
	// err5 := helper.IsExistDir("data")
	// if err5 != nil {
	// 	println(err5.Error())
	// 	return
	// }
	// println("dir exist")

	// create dir
	// err3 := helper.CreateDir("data")
	// if err3 != nil {
	// 	println(err3.Error())
	// 	return
	// }
	// println("dir created")
}

func IOExample() {
	// ioClearConsole()
	// ioOutputColor()
	ioFile()
}