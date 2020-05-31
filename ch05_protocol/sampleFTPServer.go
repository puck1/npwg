package main

import (
	"fmt"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main() {
	service := ":1202"

	addr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		// Keep alive until the client closes it or an error occurs.
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		s := string(buf[:n])
		if s[:2] == CD {
			chdir(conn, s[3:])
		} else if s[:3] == DIR {
			lsDir(conn)
		} else if s[:3] == PWD {
			pwd(conn)
		} else {
			break
		}
	}
}

func chdir(conn net.Conn, dir string) {
	err := os.Chdir(dir)
	if err == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("Error"))
	}
}

func lsDir(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	fileNames, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, fileName := range fileNames {
		conn.Write([]byte(fileName + "\r\n"))
	}
}

func pwd(conn net.Conn) {
	dir, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(dir))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
