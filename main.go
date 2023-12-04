package main

import (
	"github.com/thuongtruong1009/gouse/examples"
	"github.com/thuongtruong1009/gouse/strings"
)

func Starter() string {
	return strings.Capitalize("hello world from Gouse!\n")
}

func main() {
	// println(Starter())
	// examples.ArrayExample()
	// examples.DateExample()
	// examples.FunctionExample()
	// examples.TypeExample()
	// examples.StringExample()
	// examples.NumberExample()
	// examples.MathExample()
	// examples.IOExample()
	examples.ConsoleExample()
	// examples.HelperExample()
}
