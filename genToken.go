package main

import (
	"net"
)

func genToken() string {
	// Get Ip address as string
	conn, _ := net.Dial("ip:icmp","google.com")
	ip := conn.LocalAddr().String()
	// Surround IP with square brackets & append port
	token := "([" + ip + "]:" + "2600)"
	return token
}
