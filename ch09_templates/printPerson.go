package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name 	string
	Age 	int
	Emails	[]string
	Jobs	[]*Job
}

type Job struct {
	Employer string
	Role	 string
}

const tmpl = `The name is {{.Name}}.
The age is {{.Age}}.
{{range .Emails}}An email is {{.}}.
{{end}}{{with .Jobs}}{{range .}}An employer is {{.Employer}}
and the role is {{.Role}}.
{{end}}{{end}}`

func main() {
	job1 := Job{Employer: "Monash", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of HE"}
	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")

	// You can also use template.ParseFiles() to load template from a file.
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal(err)
	}
}
