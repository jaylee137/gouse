package tools

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/thuongtruong1009/gouse/strings"

	"github.com/thuongtruong1009/gouse/helper"
	"github.com/thuongtruong1009/gouse/io"
)

func Doc(outputPath string) {
	if len(os.Args) < 2 {
		println("Please provide at least one path to a file or directory")
		return
	}

	// clear output path
	err := os.RemoveAll(outputPath)
	if err != nil {
		fmt.Println("Error removing output path:", err)
		return
	}

	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		_, err := io.IsExistDir(path)
		if err != nil {
			println("Please provide a path to a directory")
			return
		}

		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		var functions []helper.Function
		for _, file := range files {
			if !strings.EndsWith(file.Name(), ".go") || strings.EndsWith(file.Name(), "_test.go") {
				continue
			}

			fmt.Println("Processing file:", file.Name())

			content, err := os.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			functions = append(functions, helper.ExtractFunctions(content)...)

			var result []byte

			result = append(result, []byte("# "+strings.Capitalize(strings.TrimSuffix(file.Name(), ".go"))+"\n\n")...)

			for i, function := range functions {
				if i == 0 && len(function.Import) > 0 {
					result = append(result, []byte("## Imports\n\n")...)
					result = append(result, function.HighlightImport()...)
					result = append(result, []byte("## Functions\n\n")...)
				}

				result = append(result, function.HighlightName()...)
				result = append(result, function.HighlightBody()...)
			}

			fileName := strings.Replace(file.Name(), ".go", ".md")
			subPath := filepath.Join(outputPath, path, filepath.Dir(file.Name()))

			// create file path if not exist
			err = io.CreatePath(filepath.Join(subPath, fileName))
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}

			err = os.WriteFile(filepath.Join(subPath, fileName), result, 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}

			functions = nil
			result = nil
		}
	}
}
