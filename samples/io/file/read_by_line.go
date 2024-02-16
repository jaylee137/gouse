package file

import "github.com/thuongtruong1009/gouse/io"

func ioReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
