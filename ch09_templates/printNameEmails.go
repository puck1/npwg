package main

import (
	template2 "html/template"
	"log"
	"os"
)

type Person struct {
	Name 	string
	Emails	[]string
}

const tmpl = `{{$name := .Name}}{{range .Emails}}Name is {{$name}}, email is {{.}}.
{{end}}`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
	}

	t := template2.New("Person template")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal(err)
	}
}
