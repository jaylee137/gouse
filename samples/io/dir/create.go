package dir

import "github.com/thuongtruong1009/gouse/io"

func SampleIoCreateDir() {
	err2 := io.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}
