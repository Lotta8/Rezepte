package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var recipes []Recipe

func createRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generiere eine eindeutige ID für das Rezept
	recipe.ID = len(recipes) + 1

	// Füge das Rezept zur Liste hinzu
	recipes = append(recipes, recipe)

	// Gib die erstellte Recipe-Struktur als JSON zurück
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func main() {
	http.HandleFunc("/recipes", createRecipe)

	fmt.Println("Server gestartet. Höre auf Port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
