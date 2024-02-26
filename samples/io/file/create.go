package file

import "github.com/thuongtruong1009/gouse/io"

func SampleIoCreateFile() {
	err := io.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}
