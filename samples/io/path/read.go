package path

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/io"
)

func SampleIoReadPath() {
	relativePath := "tmp/example.txt"

	content, err := io.ReadPath(relativePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}
