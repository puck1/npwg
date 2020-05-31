package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	var newTime time.Time
	// This client and server are exchanging ASN.1 encoded data values, not textual strings.
	_, err = asn1.Unmarshal(result, &newTime)
	checkError(err)

	fmt.Println("After marshal/unmarshal: ", newTime)

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetal error: %s", err.Error())
		os.Exit(1)
	}
}
