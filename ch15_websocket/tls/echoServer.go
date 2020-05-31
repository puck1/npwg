package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func Echo (conn *websocket.Conn) {
	fmt.Println("Echoing to", conn.RemoteAddr())
	defer func() {
		fmt.Println("Closing connection to", conn.RemoteAddr())
		//conn.Close()
	}()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello, %d", i)
		fmt.Println("Sending to client", conn.RemoteAddr(), ":", msg)
		err := websocket.Message.Send(conn, msg)
		//_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("Can't send")
			break
		}

		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			log.Println("Can't receive")
			break
		}
		fmt.Println("Got reply form client", conn.RemoteAddr(), ":", reply)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	// Same as https.
	// You can use `ch7_security/x509Cert/genX509Cert.go` to generate certificates.
	err := http.ListenAndServeTLS(":12345", "localhost.crt", "localhost.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
