package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	uiDir = "dir"
	uiCd = "cd"
	uiPwd = "pwd"
	uiQuit = "quit"
)

const (
	DIR = "DIR"
	CD = "CD"
	PWD = "PWD"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1] + ":" + "1202"

	addr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		// Trim space, '\t', '\r' and '\n' in the right of line.
		line = strings.TrimRight(line, " \t\r\n")
		if err != nil {
			break
		}

		// Get first 2 arguments(if has).
		strs := strings.SplitN(line, " ", 2)
		switch strs[0] {
		case uiDir:
			dirRequest(conn)
		case uiCd:
			cdRequest(conn, strs[1])
		case uiPwd:
			pwdRequest(conn)
		case uiQuit:
			conn.Close()
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "Unknown command")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte("DIR" + " "))

	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, _ := conn.Read(buf[0:])
		result.Write(buf[:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[ : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))

	var response [512]byte
	n, _ := conn.Read(response[0:])
	if s := string(response[:n]); s != "OK" {
		fmt.Fprintln(os.Stderr, "Failed to change directory")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD + " "))

	var response [512]byte
	n ,_ := conn.Read(response[0:])
	s := string(response[:n])
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
