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
	isDev := flag.Bool("isDev", false, "a bool type for toggle enviroment mode")
	flag.Parse()
	if *isDev {
		samples.Run()
	}
}
