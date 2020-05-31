package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name 	Name
	Email 	[]Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for i := 0; i < 10; i++ {
			var person Person
			err = decoder.Decode(&person)
			if err != nil {
				break
			}

			err = encoder.Encode(person)
			if err != nil {
				break
			}
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
