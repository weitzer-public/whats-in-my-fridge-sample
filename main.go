package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
)

// main is the entry point for the application.
// It sets up the HTTP handlers and starts the web server.
func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, config.GCPProjectID, "us-central1")
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	server := &Server{
		geminiClient: client,
		config:       config,
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/suggest-recipe", server.suggestRecipeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}