package main

import (
	"encoding/gob"
	"fmt"
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
	var person Person
	loadGob("person.gob", &person)
	fmt.Println("Person", person)
}

func loadGob(file string, key interface{}) {
	fd, err := os.Open(file)
	checkError(err)

	decoder := gob.NewDecoder(fd)
	err = decoder.Decode(key)
	checkError(err)

	fd.Close()
}

func checkError(err error)  {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
