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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	// In general, we get []byte from wire then transfer it to []uint16 with BOM.
	// Then we need to decode it into runes and cast it to a string.
	shorts := readShorts(conn)
	runes := utf16.Decode(shorts)
	fmt.Println(string(runes))
}

func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte
	n, err := conn.Read(buf[:2])
	for {
		m ,err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
	}
	checkError(err)

	var shorts []uint16
	shorts = make([]uint16, n/2)
	if buf[0] == 0xff && buf[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2 - 1] = uint16(buf[i]) << 8 + uint16(buf[i+1])
		}
	} else if buf[0] == 0xfe && buf[1] == 0xff {
		// small endian
		for i := 2; i < n; i += 2 {
			shorts[i/2 - 1] = uint16(buf[i]) + uint16(buf[i+1]) << 8
		}
	} else {
		fmt.Fprintln(os.Stderr, "Unknown order")
	}
	return shorts
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
