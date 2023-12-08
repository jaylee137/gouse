package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/thuongtruong1009/gouse/helper"
	"github.com/thuongtruong1009/gouse/io"
)

func main() {
	var outputPath = "docs/gen"

	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one path to a file or directory")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		_, err := io.IsExistDir(path)
		if err != nil {
			fmt.Println("Please provide a path to a directory")
			return
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		var functions []helper.Function
		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".go") {
				continue
			}

			content, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			functions = append(functions, helper.ExtractFunctions(content)...)

			var result []byte
			for _, function := range functions {
				result = append(result, function.HighlightName()...)
				result = append(result, function.HighlightImport()...)
				result = append(result, function.HighlightBody()...)
			}

			fileName := strings.TrimSuffix(file.Name(), ".go") + ".md"
			subPath := filepath.Join(outputPath, path, filepath.Dir(file.Name()))

			// create file path if not exist
			err = io.CreatePath(filepath.Join(subPath, fileName))
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}

			err = ioutil.WriteFile(filepath.Join(subPath, fileName), result, 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}
		}
	}
}
