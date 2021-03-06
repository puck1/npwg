package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "http://proxy-host:port", "http://host:port/page")
		os.Exit(1)
	}
	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	checkError(err)
	rawString := os.Args[2]
	rawURL, err := url.Parse(rawString)
	checkError(err)

	// You can specify a proxy url like above or
	// get it from environment such as HTTP_PROXY or http_proxy for A SPECIFIC REQUEST, like below.
	// proxyURL, err = http.ProxyFromEnvironment(request)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),		// Set proxyURL.
	}

	client := http.Client{					// NOTE proxy is set in Transport field in a client object but not in a request.
		Transport:     transport,			// It is a pointer to a http.Transport object.
											// You can also set tls config(rootCA, client cert) in transport of client.
		Timeout:       5 * 1000 * 1000 * 1000,
	}

	request, err := http.NewRequest("GET", rawURL.String(), nil)

	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println("[Request Header]")
	fmt.Println(string(dump))

	response, err := client.Do(request)
	checkError(err)
	fmt.Println("**Read ok**")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("**Response ok**")

	dump, _ = httputil.DumpResponse(response, false)
	fmt.Println("[Response Header]")
	fmt.Println(string(dump))

	var buf [512]byte
	reader := response.Body
	fmt.Println("[Response Body]")
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
	os.Exit(1)				// Never reached.
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
