package main

import (
	"fmt"
	// "github.com/thuongtruong1009/gouse/examples"
	"github.com/thuongtruong1009/gouse/strings"
)

func Starter() string {
	return strings.Capitalize("hello world!")
}

func main() {
	fmt.Println(Starter())
	
	// examples.ArrayExample()
	// examples.DateExample()
	// examples.UtilExample()
	// examples.TypeExample()
}