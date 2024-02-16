package file

import "github.com/thuongtruong1009/gouse/io"

func ioTruncateFile() {
	err := io.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}
