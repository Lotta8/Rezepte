package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type Recipe struct {
	ID          int      `json:"id" :"id"`
	Name        string   `json:"name" :"name"`
	Ingredients []string `json:"ingredients" :"ingredients"`
}

var recipes = []Recipe{
	{ID: 1, Name: "Würziger Fleischsalat", Ingredients: []string{"Fleischwurst", "Schinkenwurst", "Essiggurken", "Mayonnaise", "Miracle Whip", " Senf", "Pfeffer", "Salz", "Zucker"}},
	{ID: 2, Name: "Griessklösschen", Ingredients: []string{"Butter", "Milch", "Hartweisengriess", "Petersilie", "Salz", "Pfeffer", "Eier"}},
	{ID: 3, Name: "Spargelcremesuppe", Ingredients: []string{"Weisse Spargeln", "Butter", "Schalotten", "Zucker", "Mehl", "Gemüsebouillon", "Rahm", "Salz", "Pfeffer"}},
}

func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(recipes)
}

func main() {
	http.HandleFunc("/recipes/all", getAllRecipes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
