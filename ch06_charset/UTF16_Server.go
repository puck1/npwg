package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

// "\u"(or U+): unicode
// NOTE: It's a character of two bytes.
const BOM = '\ufffe'
//const BOM = 0xfffe 	// This is a int const.

func main() {
	service := "0.0.0.0:1210"
	addr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		str := "j'ai arrêté"
		// utf16.Encode encodes a slice of rune into a slice of uint16
		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)
		conn.Close() // we're finished
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte
	bytes[0] = BOM >> 8
	//bytes[1] = BOM & 255
	bytes[1] = BOM & '\u00ff'

	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}

	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		_, err = conn.Write(bytes[0:])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
