package file

import "github.com/thuongtruong1009/gouse/io"

func ioRenameFile() {
	err := io.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}
