package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../x509Cert/localhost.pem", "../x509Cert/private.pem")
	checkError(err)

	config := tls.Config{ Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time {
		return now
	}
	config.Rand = rand.Reader

	service := ":1200"
	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	defer listener.Close()
	fmt.Println("[ Listening on", listener.Addr(), "]")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("[ Accepted connection from", conn.RemoteAddr(), "]")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("[ Closed connection from", conn.RemoteAddr(), "]")
	}()

	var buf [512]byte
	for {
		//fmt.Println("[ Trying to read from", conn.RemoteAddr(), "]")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("[", err, "]")
			return
		}
		fmt.Println("[ Read:", string(buf[:n]), "]")
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("[", err, "]")
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

