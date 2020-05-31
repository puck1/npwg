package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string to hash:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()							// Scan read a line and trim '\n' by default.
	s := scanner.Text()
	s = strings.TrimRight(s, " \t\r")

	fmt.Printf("Writing MD5 hash of \"%s\"...\n", s)

	hash := md5.New()
	_, err := io.WriteString(hash, s)		// Remember common Read or Write func can only manipulate []byte.
	//hash.Write([]byte(s))					// Same.
	checkError(err)

	hashValue := hash.Sum(nil)
	hashSize := hash.Size()					// 16

	fmt.Print("Result: ")
	for i := 0; i < hashSize; i += 4 {
		var val uint32
		val = uint32(hashValue[i]) << 24 +
			uint32(hashValue[i+1]) << 16 +
			uint32(hashValue[i+2]) << 8 +
			uint32(hashValue[i+3])
		fmt.Printf("%x ", val)					// Print out four bytes in hexadecimal format for four times.
		//fmt.Printf("%x ", hashValue[i:i+4])	// Same.
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
