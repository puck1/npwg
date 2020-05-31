package main

import (
	"./xmlcodec"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

type Person struct {
	Name string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "ws://host:port/page")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := websocket.Dial(service,"", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	var person Person
	err = xmlcodec.XMLCodec.Receive(conn, &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", person.Name)
	for _, v := range person.Emails {
		fmt.Println("An email:", v)
	}
	conn.Close()
}


