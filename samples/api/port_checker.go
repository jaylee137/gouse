package api

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/api"
)

/*
Title: Port Checker
Description: Check if a port is open
Package: api
Input: protocol, hostname, port
*/

func SampleApiPortChecker() {
	open := api.PortChecker("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
