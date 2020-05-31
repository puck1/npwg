package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	// The only difference to use websocket client.
	conn, err := websocket.Dial(service, "", "http://localhost")
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
