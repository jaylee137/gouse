package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/shared"
	"github.com/thuongtruong1009/gouse/io"
)

func ioCheckDir() {
	isExist, err1 := io.IsExistDir("tmp")
	if err1 != nil {
		println(err1.Error())
	}
	if isExist {
		println("dir exist")
	} else {
		println("dir not exist")
	}
}

func ioCreateDir() {
	err2 := io.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}

func ioRemoveDir() {
	err3 := io.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}

func ioLsDir() {
	data, err := io.LsDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func ioHierarchyDir() {
	data, err := io.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}

func ioCurrentDir() {
	data, err := io.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
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
	_, err4 := io.IsExistFile(exampleFile)
	if err4 != nil {
		println(err4.Error())
		return
	}
	println("file exist")

	// read file
	data, err := io.ReadFileObj[*User](exampleFile)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("data: %+v\n", data)

	// write file
	err2 := io.WriteFileObj(exampleFile, exampleData)
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

func ioCheckFile() {
	isExist, err := io.IsExistFile("data.json")
	if err != nil {
		println(err.Error())
	}
	if isExist {
		println("file exist")
	} else {
		println("file not exist")
	}
}

func ioCreateFile() {
	err := io.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}

func ioRemoveFile() {
	err := io.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}

func ioReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}

func ioFileInfo() {
	data, err := io.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	fmt.Println("Name of the File:", data.Name())
    fmt.Println("Size of File:", data.Size())
    fmt.Println("Permissions File:", data.Mode())
    fmt.Println("Last Modified of File:", data.ModTime())
    fmt.Println("Directory check: ", data.IsDir())
    fmt.Printf("System Process info: %+v\n\n", data.Sys())
}

func ioRenameFile() {
	err := io.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}

func IOExample() {	
	// ioCheckDir()
	// ioCreateDir()
	// ioRemoveDir()
	// ioLsDir()
	// ioHierarchyDir()
	// ioCurrentDir()
	
	// ioFile()
	// ioCheckFile()
	// ioCreateFile()
	// ioRemoveFile()
	// ioReadFileByLine()
	// ioFileInfo()
	// ioRenameFile()

	// ioClearConsole()
	// ioOutputColor()
}