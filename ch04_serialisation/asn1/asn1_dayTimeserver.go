/*
 * ASN.1 also includes some "useful types" such as UTC time.
 * Go supports this UTC time type. This means that you can pass time values in a way
 * that is not possible for other data values.
 * ASN.1 does not support pointers, but Go has special code to manage pointers to time values.
 * The function GetLocalTime returns `*time.Time`.
 * The special code marshals this, and it can be unmarshalled into a pointer variable to a `time.Time` object.
*/
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		dayTime := time.Now()
		// Ignore return error.
		mdata, _ := asn1.Marshal(dayTime)
		conn.Write(mdata)
		conn.Close()	// We are finished.
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetal error: %s", err.Error())
		os.Exit(1)
	}
}
