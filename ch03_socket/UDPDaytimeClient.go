package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	// Transfer string("host:port" or "ip:port") to a UDPAddr instance.
	// Parameter network is "udp4""udp6" or "udp",
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	// Client uses DialUDP to create a connection.
	// Parameter laddr is local address and raddr is remote address.
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1) }
}