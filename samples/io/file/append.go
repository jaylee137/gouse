package file

import "github.com/thuongtruong1009/gouse/io"

func ioAppendToFile() {
	err := io.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}
