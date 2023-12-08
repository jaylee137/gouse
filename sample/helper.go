package sample

import (
	"github.com/thuongtruong1009/gouse/helper"
)

func helperRandomID() {
	randomID := helper.RandomID()
	println("randomID: ", randomID)
}

func helperAutoMdDoc() {
	inputFilePath := "helper/doc.go"
	outputFilePath := "helper/doc.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}

func HelperExample() {
	// helperRandomID()
	helperAutoMdDoc()
}
