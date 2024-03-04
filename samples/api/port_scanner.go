package api

import "github.com/thuongtruong1009/gouse/api"

/*
Title: Port Scanner
Description: This sample will scan for open ports on a given host.
Package: api
Input: protocol, hostname, start, end
*/

func SampleApiPortScanner() {
	api.PortScanner("tcp", "127.0.0.1", 3000, 8080)
}
