package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/net"
	"github.com/thuongtruong1009/gouse/types"
)

func netOpen() {
	net.Open("https://google.com")
}

func netEncode() {
	println("Encode: ", net.Encode("https://google.com"))
}

func netDecode() {
	println("Decode: ", net.Decode("https%3A%2F%2Fgoogle.com"))
}

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

func netHeader() {
	header, err := net.Header("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(types.MapAsString(header))
}

func netConnectTime() {
	connectTime, err := net.ConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
