package api

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func PortScanner(protocol, hostname string, start, end int) {
	for i := start; i < end; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial(protocol, hostname+":"+port)
		if err == nil {
			fmt.Println("Port", i, "open")
			conn.Close()
		} else {
			fmt.Println("Port", i, "closed")
		}
	}
}

func PortChecker(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	return true
}
