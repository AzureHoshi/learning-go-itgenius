package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define the port on which the server will run
const port = 8080

type application struct {
	domain string
}

func main() {
	// set application configuration
	var app application
	app.domain = "localhost"

	// read from command line arguments

	// connect to database

	// start the server
	log.Printf("Starting server on %s:%d", app.domain, port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
