package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8} 		// encoding src
	bb := bytes.Buffer{}								// encoding dsc
	// We can transfer binary bytes to ASCII form so that it is easy to read and check.

	// NewEncoder creates an base64 encoder by specifying encoding method and the io.Writer to write.
	encoder := base64.NewEncoder(base64.StdEncoding, &bb)		// Note it is &bytes.Buffer that implements io.Writer.
	// Write writes its parameter into the io.Writer specified in NewEncoder.
	encoder.Write(eightBitData)
	// encoder is a implementation of io.WriteCloser.
	encoder.Close()

	// Binary to text.
	fmt.Println(&bb) 				// Note &bytes.Buffer implements Stringer interface, so this just print out a string.

	dbuf := make([]byte, 12)		// NOTE: 8 will be wrong.
	// NewDecoder creates an base64 decoder by specifying encoding method and the io.Reader to read from.
	decoder := base64.NewDecoder(base64.StdEncoding, &bb)
	// Write reads from the io.Reader specified in NewDecoder to its parameter byte slice.
	decoder.Read(dbuf)
	for _, ch := range dbuf {
		fmt.Print(ch)
	}
}


