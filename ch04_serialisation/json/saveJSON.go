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

func main() {
	person := Person{
		Name:  Name{ Family:   "Du", Personal: "Yaodi"},
		Email: []Email{{Kind: "domestic", Address: "domestic@server.com"},
						{Kind: "foreign", Address: "foreign@server.com"}},
	}

	saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {
	fd, err := os.Create(fileName)		// Even if file exists, os.Create won't return an error.
	checkError(err)
	encoder := json.NewEncoder(fd)
	err = encoder.Encode(key)
	checkError(err)
	fd.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
