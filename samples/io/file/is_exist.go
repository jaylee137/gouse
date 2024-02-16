package file

import "github.com/thuongtruong1009/gouse/io"

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
