package helper

import "github.com/thuongtruong1009/gouse/helper"

func helperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}
