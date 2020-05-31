package main

import (
	"./xmlcodec"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Emails []string
}

func sendPerson(conn *websocket.Conn) {
	defer conn.Close()
	person := Person{
		Name: "Puck",
		Emails: []string{"test@offical.com", "test@home.com"},
	}
	err := xmlcodec.XMLCodec.Send(conn, person)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.Handle("/", websocket.Handler(sendPerson))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal(err)
	}
}


