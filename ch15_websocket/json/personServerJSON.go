package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	"net/http"
)

type Person struct {
	Name 	string
	Emails []string
}

func receiveJSON (conn *websocket.Conn) {
	fmt.Println("Trying to receive json from", conn.RemoteAddr())

	var person Person
	for {
		err := websocket.JSON.Receive(conn, &person)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client:", conn.RemoteAddr())
				break
			}
			log.Println(err)
			break
		} else {
			fmt.Println("Got Person")
			fmt.Println("Name:", person.Name)
			fmt.Print("Emails: ")
			for _, email := range person.Emails {
				fmt.Print(email + " ")
			}
			fmt.Println()
		}
	}
}

func sendJSON (conn *websocket.Conn) {
	fmt.Println("Sending json to", conn.RemoteAddr())
	defer func() {
		fmt.Println("Closing connection to", conn.RemoteAddr())
		// Will close conn automatically when handler returns.
		//conn.Close()
	}()

	// Send twice.
	person := Person{
		Name:   "Puck",
		Emails: []string{"fake@gmail.com", "fake@outlook.com"},
	}
	err := websocket.JSON.Send(conn, person)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Send Person successfully")
		fmt.Println("Name:", person.Name)
		fmt.Print("Emails: ")
		for _, email := range person.Emails {
			fmt.Print(email + " ")
		}
		fmt.Println()
	}

	person = Person{
		Name:   "Puck1",
		Emails: []string{"fake1@gmail.com", "fake1@outlook.com"},
	}
	err = websocket.JSON.Send(conn, person)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Send Person successfully")
		fmt.Println("Name:", person.Name)
		fmt.Print("Emails: ")
		for _, email := range person.Emails {
			fmt.Print(email + " ")
		}
		fmt.Println()
	}
}

func main() {
	// Register websocket handler functions.
	http.Handle("/sendJSON", websocket.Handler(receiveJSON))
	http.Handle("/getJSON", websocket.Handler(sendJSON))

	//http.ListenAndServe(":12345", nil)

	// Listen tcp.
	service := ":12345"
	l, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening at", l.Addr())

	// Serve.
	err = http.Serve(l, nil)		// Use default mux.
	if err != nil {
		log.Fatal(err)
	}
}
