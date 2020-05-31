package main

import (
	"log"
	"net/http"
)

func main() {
	// Param is a local dir and it will be the file server root dir.
	// You can use "get.go" or "clientGet.go" in "../client" dir as client, or just use a browser.
	fileServer := http.FileServer(http.Dir("./httpd/db/"))		// http.Dir is a type cast.

	// Register the handler and deliver requests to it.
	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)
	// That's it!
	// This server even delivers "404 not found" messages for requests for file resources that don't exist!
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
