package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ShoppingCart struct {
	Recipes []Recipe `json:"recipes"`
}

var shoppingCart ShoppingCart

func deleteRecipeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	for index, recipe := range shoppingCart.Recipes {
		if recipe.ID == id {
			// Entferne das Rezept aus dem Warenkorb
			shoppingCart.Recipes = append(shoppingCart.Recipes[:index], shoppingCart.Recipes[index+1:]...)
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}

	http.Error(w, "Recipe not found in shopping cart", http.StatusNotFound)
}

func SetupRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/shoppingcart/{id}", deleteRecipeByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
