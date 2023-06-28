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
In diesem Beispiel wird ein HTTP-Server gestartet, der POST-Anfragen auf den Pfad "/recipes" akzeptiert. Wenn eine POST-Anfrage empfangen wird, wird die Funktion createRecipe aufgerufen. Die Funktion liest den JSON-Körper der Anfrage, dekodiert ihn in eine Recipe-Struktur und fügt das Rezept zur Liste recipes hinzu. Anschließend wird das erstellte Rezept als JSON in der Antwort zurückgegeben.

Du kannst diesen Code verwenden, um einen einfachen HTTP-Server in Go zu starten, der POST-Anfragen zum Erstellen von Rezepten akzeptiert. Bitte beachte, dass dies nur ein grundlegendes Beispiel ist und nicht die vollständige Validierung oder Speicherung der Rezepte enthält. Du kannst den Code entsprechend anpassen, um deine spezifischen Anforderungen zu erfüllen.






