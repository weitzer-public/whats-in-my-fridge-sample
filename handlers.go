package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
)

// Server holds the Gemini client.
type Server struct {
	geminiClient *genai.Client
	config       *Config
}

// PageData holds all the data for the template.
type PageData struct {
	Recipe              template.HTML
	OriginalIngredients string
}

// rootHandler renders the main page.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, PageData{}); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// suggestRecipeHandler handles the recipe suggestion request.
func (s *Server) suggestRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	ingredients := r.FormValue("ingredients")

	prompt := fmt.Sprintf("Generate a simple recipe using primarily these ingredients: %s. Include a title, a bulleted list of ingredients, and numbered instructions.", ingredients)

	model := s.geminiClient.GenerativeModel(s.config.GeminiModel)
	resp, err := model.GenerateContent(r.Context(), genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	recipe := ""
	if len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				recipe += string(txt)
			}
		}
	}

	recipeHTML := template.HTML(recipe)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Recipe:              recipeHTML,
		OriginalIngredients: ingredients,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
