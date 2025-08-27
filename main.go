package main

import (
	"log"
	"net/http"
)

// main is the entry point for the application.
// It sets up the HTTP handlers and starts the web server.
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/suggest-recipe", suggestRecipeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
