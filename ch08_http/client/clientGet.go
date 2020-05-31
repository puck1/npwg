package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)

	// You can modify client behavior such as setting timeout by this means.
	client := http.Client{
		Timeout: 5 * 1000 * 1000 * 1000,		// A 5-second timeout.
	}
	//client := http.DefaultClient

	// You can modify request by this means.
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)
	// Only accept UTF-8.
	request.Header.Add("Accept-CharSet", "UTF-8;q=1, ISO-8859-1;q=0")

	response, err := client.Do(request)
	checkError(err)

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("[Response header]")
	b, _ := httputil.DumpResponse(response, false)
	// httputil.DumpResponse transfers a response to bytes.
	// You can use http.ReadResponse to recover it into a response object.
	fmt.Println(string(b))

	chSet := getChatSet(response)
	fmt.Printf("[Charset: %s]\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("[response body]")
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
		fmt.Print(string(buf[:n]))
	}
	os.Exit(1)		// Never reached.
}

func getChatSet(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
