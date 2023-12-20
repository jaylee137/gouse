package sample

import (
	"github.com/thuongtruong1009/gouse/helper"
)

func helperRandomID() {
	println("Generate random ID: ", helper.RandomID())
}

func helperUUID() {
	println("New uuid: ", helper.UUID())
}

func helperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}
