package net

import "github.com/thuongtruong1009/gouse/net"

func netCheck() {
	ok, err := net.Check("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}

func netCheckWithStatusCode() {
	statusCode, err := net.CheckWithStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}
