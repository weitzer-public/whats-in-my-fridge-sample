package main

import (
	"html/template"
	"log"
	"net/http"
)

// PageData holds all the data for the template.
type PageData struct {
	Recipe            template.HTML
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
func suggestRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	ingredients := r.FormValue("ingredients")
	recipe := getRecipeSuggestion(ingredients)

	// Intentionally introduce XSS vulnerability by converting the string to template.HTML
	// This bypasses Go's automatic HTML escaping.
	recipeHTML := template.HTML(recipe)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Recipe:            recipeHTML,
		OriginalIngredients: ingredients,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// getRecipeSuggestion is a placeholder function that returns a hardcoded recipe.
// It intentionally uses the raw user input to construct the prompt, creating a prompt injection vulnerability.
func getRecipeSuggestion(ingredients string) string {
	// Prompt injection vulnerability: The user's input is directly concatenated into the "prompt".
	// A malicious user could inject commands or other text.
	prompt := "Suggest a simple recipe based on the following ingredients: " + ingredients
	log.Printf("Generated prompt: %s", prompt)

	// Hardcoded recipe for demonstration.
	// The recipe includes HTML tags that will be rendered due to the XSS vulnerability.
	return "<h2>Simple Pasta Aglio e Olio</h2>\n" +
		"<p>A classic Italian dish that's quick and delicious.</p>\n" +
		"<h3>Ingredients:</h3>\n" +
		"<ul>\n" +
		"  <li>200g spaghetti</li>\n" +
		"  <li>4 cloves garlic, thinly sliced</li>\n" +
		"  <li>1/2 cup olive oil</li>\n" +
		"  <li>1/4 teaspoon red pepper flakes</li>\n" +
		"  <li>Salt and black pepper to taste</li>\n" +
		"  <li>2 tablespoons fresh parsley, chopped</li>\n" +
		"</ul>\n" +
		"<h3>Instructions:</h3>\n" +
		"<ol>\n" +
		"  <li>Cook spaghetti according to package directions.</li>\n" +
		"  <li>While pasta is cooking, heat olive oil in a large skillet over medium heat. Add garlic and red pepper flakes and cook until garlic is golden brown.</li>\n" +
		"  <li>Drain pasta and add it to the skillet. Toss to combine.</li>\n" +
		"  <li>Season with salt and pepper, and stir in parsley.</li>\n" +
		"</ol>"
}
