package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type RecipeOne struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(recipes)
}

func main() {
	http.HandleFunc("/recipes/all", getAllRecipes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
