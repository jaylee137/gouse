package helper

import "github.com/thuongtruong1009/gouse/helper"

func SampleHelperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}
