package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

// Websocket handlers are different from http handlers,
// the parameter is a pointer to websocket.Conn and both
// client and server can write to or read from it as long as it is not closed.
func Echo (conn *websocket.Conn) {
	fmt.Println("Echoing to", conn.RemoteAddr())
	defer func() {
		fmt.Println("Closing connection to", conn.RemoteAddr())
		// Will close conn automatically when handler returns.
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
	// It is easy to handle websocket, you just need to register a websocket handler to http.defaultMux or your own mux.
	http.Handle("/", websocket.Handler(Echo))			// Register handler.
	err := http.ListenAndServe(":12345", nil)		// Use default mux.
	if err != nil {
		log.Fatal(err)
	}

	// Use a complete websocket.Server:
	//s := websocket.Server{
	//	Config:    websocket.Config{},
	//	Handshake: nil,
	//	Handler:   Echo,
	//}
	//err := http.ListenAndServe(":12345", s)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
