package sample

import (
	"github.com/thuongtruong1009/gouse/helper"
)

func helperRandomID() {
	println("randomID: ", helper.RandomID())
}

func helperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}

func HelperExample() {
	helperRandomID()
	helperAutoMdDoc()
}