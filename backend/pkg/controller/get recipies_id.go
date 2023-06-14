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

var recipe = []Recipe{
	{ID: 1, Name: "Würziger Fleischsalat", Ingredients: []string{"Fleischwurst", "Schinkenwurst", "Essiggurken", "Mayonnaise", "Miracle Whip", " Senf", "Pfeffer", "Salz", "Zucker"}},
	{ID: 2, Name: "Griessklösschen", Ingredients: []string{"Butter", "Milch", "Hartweisengriess", "Petersilie", "Salz", "Pfeffer", "Eier"}},
	{ID: 3, Name: "Spargelcremesuppe", Ingredients: []string{"Weisse Spargeln", "Butter", "Schalotten", "Zucker", "Mehl", "Gemüsebouillon", "Rahm", "Salz", "Pfeffer"}},
}

func getRecipeByID(
	w http.ResponseWriter,
	r *http.Request,
) {
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
