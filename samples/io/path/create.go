package path

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/io"
)

func SampleIoCreatePath() {
	relativePath := "tmp/example.txt"

	if err := io.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}
