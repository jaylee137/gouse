package dir

import "github.com/thuongtruong1009/gouse/io"

func SampleIoRemoveDir() {
	err3 := io.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}
