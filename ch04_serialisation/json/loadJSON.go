package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name 	Name 		`json:"personName"`
	Email 	[]Email		`json:"personEmail"`
}

type Name struct {
	Family string		`json:"familyName"`
	Personal string		`json:"personalName"`
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
	loadJSON("person.json", &person)
	fmt.Println("Person", person)
}

func loadJSON(fileName string, key interface{}) {
	fd, err := os.Open(fileName)
	checkError(err)

	decoder := json.NewDecoder(fd)

	err = decoder.Decode(key)
	checkError(err)
	fd.Close()			// Don't forget to close file after used!
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}