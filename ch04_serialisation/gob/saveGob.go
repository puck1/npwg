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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveGob("person.gob", person)
}

func saveGob(fileName string, key interface{}) {
	fd ,err := os.Create(fileName)		// Even if file exists, os.Create won't return an error.
	checkError(err)
	encoder := gob.NewEncoder(fd)
	err = encoder.Encode(key)
	checkError(err)
	fd.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetal error: ", err)
		os.Exit(1)
	}
}
