package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "url")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Get(url)		// Same as http.DefaultClient.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("[Response header]")
	b, _ := httputil.DumpResponse(response, false)
	// httputil.DumpResponse transfers a response to bytes.
	// You can use http.ReadResponse to recover it into a response object.
	fmt.Println(string(b))

	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("[Response body]")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			if err == io.EOF && n > 0 {
				// It's strange that once resp.Body reach EOF then it will return EOF, not next time.
				fmt.Print(string(buf[:n]))
			}
			reader.Close()			// The body object is an io.ReadCloser.
			os.Exit(0)
		}
		fmt.Println(string(buf[:n]))
	}
	os.Exit(2)				// Never reached.
}

func acceptableCharset(contentTypes []string) bool {
	// each type is like [text/html; charset=UTF-8]
	// we want the UTF-8 only
	// NOTE this function just seek strings for convenience but some response does not contain a charset.
	for _, cType := range contentTypes {
		if strings.Index(strings.ToUpper(cType), "UTF-8") != -1 {
			return true
		}
	}
	return false
}
