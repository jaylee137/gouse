package file

import "github.com/thuongtruong1009/gouse/io"

func ioCopyFile() {
	err := io.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}
