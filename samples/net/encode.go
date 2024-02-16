package net

import "github.com/thuongtruong1009/gouse/net"

func netEncode() {
	println("Encode: ", net.Encode("https://google.com"))
}

func netDecode() {
	println("Decode: ", net.Decode("https%3A%2F%2Fgoogle.com"))
}
