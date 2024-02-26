package main

import (
	"flag"

	"github.com/thuongtruong1009/gouse/samples"
	"github.com/thuongtruong1009/gouse/strings"
)

func Starter() {
	println(strings.Capitalize("hello world from Gouse!"))
}

func main() {
	Starter()

	// only for development mode testing
	isDevFlag := flag.Bool("isDev", false, "toggle enviroment mode")
	flag.Parse()
	if *isDevFlag {
		samples.Run()
	}
}
