package path

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/io"
)

func ioZip() {
	filesToZip := []string{"file1.txt", "file2.txt"}
	zipFileName := "archive.zip"
	err := io.Zip(zipFileName, filesToZip)
	if err != nil {
		fmt.Println("Error zipping files:", err)
	}

	println("Files zipped successfully:", zipFileName)
}
