package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Note http.HandlerFunc is a type.
	myHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Just return no content - arbitrary headers can be set, arbitrary body.
		fmt.Println("Request: " + request.URL.String())
		writer.WriteHeader(http.StatusNoContent)
	})

	err := http.ListenAndServe(":8001", myHandler)
	if err != nil {
		log.Fatal(err)
	}
}
