package samples

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/api"
)

func apiPortScanner() {
	fmt.Println("Port Scanning")
	// <protocol>, <hostname>, <start port>, <end port>
	api.PortScanner("tcp", "127.0.0.1", 3000, 8080)
}

func apiPortChecker() {
	fmt.Println("Port Checker")
	// <protocol>, <hostname>, <port>
	open := api.PortChecker("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
