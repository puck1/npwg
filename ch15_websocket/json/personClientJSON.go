package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	url2 "net/url"
	"os"
)

type Person struct {
	Name string
	Emails []string
}

func sendJSONTo(conn *websocket.Conn) {
	fmt.Println("Trying to send json")
	defer func() {
		fmt.Println("Closing connection to remote")
		// Close connection by client.
		conn.Close()
	}()

	// Send twice.
	person := Person{
		Name:   "Puckpu",
		Emails: []string{"test@offical.com", "test@school.edu.com", "TEST@home.com"},
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
		Name:   "Puckpu1",
		Emails: []string{"test1@offical.com", "test1@school.edu.com", "TEST1@home.com"},
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

func getJSONFrom(conn *websocket.Conn) {
	fmt.Println("Trying to receive json")
	var person Person
	for {
		err := websocket.JSON.Receive(conn, &person)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by remote")
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

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "ws://host:port/page")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to %v successfully\n", conn.RemoteAddr())

	url, err := url2.Parse(service)
	if err != nil {
		log.Fatal(err)
	}

	switch url.EscapedPath() {
	case "/sendJSON":
		sendJSONTo(conn)
	case "/getJSON":
		getJSONFrom(conn)
	}

	os.Exit(0)
}
