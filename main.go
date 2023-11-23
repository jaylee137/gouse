package main

import (
	"github.com/thuongtruong1009/gouse/examples"
	"github.com/thuongtruong1009/gouse/strings"
)

func Starter() string {
	return strings.Capitalize("hello world!")
}

func main() {	
	// examples.ArrayExample()
	// examples.DateExample()
	// examples.UtilExample()
	// examples.TypeExample()
	examples.StringExample()
}