package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name 	string
	Emails	[]string
}

const tmpl = `The name is {{.Name}}.
{{range .Emails}}An email is {{. | emailExpand}}.
{{end}}`

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)		// type assert
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	indexDot := strings.LastIndexByte(substrs[1], '.')
	if indexDot == -1 || indexDot == len(substrs[1]) {
		return s
	}

	//return substrs[0] + " [AT] " + substrs[1][:indexDot] + " [DOT] " + substrs[1][indexDot+1:]
	return strings.ReplaceAll(substrs[0] + " [AT] " + substrs[1], ".", " [DOT] ")
}

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
	}

	t := template.New("Person template")

	// Add our function.
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal(err)
	}
}
