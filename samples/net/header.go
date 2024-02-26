package net

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/net"
	"github.com/thuongtruong1009/gouse/types"
)

func SampleNetHeader() {
	header, err := net.Header("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(types.MapAsString(header))
}
