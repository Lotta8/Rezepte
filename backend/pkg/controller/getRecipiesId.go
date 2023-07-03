package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Recip struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

var recipes []Recip // Datenquelle für Rezepte

func getRecipeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		if recipe.ID == id {
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}

	http.Error(w, "Recipe not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/recipes/{id}", getRecipeByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
