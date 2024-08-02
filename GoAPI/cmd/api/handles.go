package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for handling the root route here

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "success",
		Message: "Welcome to the Go API",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) HandleHealth(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for handling the health route here

	w.Write([]byte("Server is healthy"))
}
