package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// File handler for most files.
	fileServer := http.FileServer(http.Dir("./httpd/db/"))
	// param of Handle is a http.HandlerFunc
	http.Handle("/", fileServer)

	// Function handler for /cgi-bin/printenv.
	// param of HandleFunc is a common func
	http.HandleFunc("/cgi/bin/printenv", printEnv)

	// Handle and HandleFunc add handlers to DefaultServeMux.

	// Deliver requests to the handlers.
	// NOTE the param "handler" of this func must be nil to dispatch calls to different handlers above.
	// Reason: ListenAndServe will use http.DefaultServeMux if you pass nil as the second parameter.
	// For more detail of multiplexer, see https://stackoverflow.com/questions/40478027/what-is-an-http-request-multiplexer
	// and https://www.cnblogs.com/yjf512/archive/2012/08/22/2650873.html .
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
	// That's it!
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<!-- Note: for simplicity this program does not deliver well-formed HTML." +
		"It is missing html, head and body tags.-->\n<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
