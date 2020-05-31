package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "wss://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	rootCAFile := "./localhost.crt"
	rootCABytes, err := ioutil.ReadFile(rootCAFile)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	// crt file can be decoded by this function as well.
	if !certPool.AppendCertsFromPEM(rootCABytes) {
		log.Fatal("Cannot get rootCA")
	}

	tlsConfig := tls.Config{
		RootCAs: certPool,
	}
	wsConfig, err := websocket.NewConfig(service, "https://localhost")
	if err != nil {
		log.Fatal(err)
	}
	wsConfig.TlsConfig = &tlsConfig

	conn, err := websocket.DialConfig(wsConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to %v successfully\n", conn.RemoteAddr())

	var msg string
	for {
		err = websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				fmt.Println("Connection closed by remote")
				break
			}
			log.Println(err)
			break
		}
		fmt.Println("Received from remote:", msg)

		err = websocket.Message.Send(conn, msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
	os.Exit(0)
}
