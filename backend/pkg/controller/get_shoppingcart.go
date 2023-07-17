package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShoppingCartItemsResponse struct {
	Recipes     []RecipeOverviewResponse `json:"recipe"`
	Ingredinets []IngredientResponse     `json:"ingredinets"`
}

type IngredientResponse struct {
	Menge       float64 `json:"menge"`
	Bezeichnung string  `json:"bezeichnung"`
}

func (h *Handler) GetShoppingCart(context *gin.Context) {
	shoppingCartItems := h.shoppingCardInMemoryStorage.Get()
	// Berechnung des Response Objekt ...
	context.JSON(http.StatusOK, shoppingCartItems)
}

//type ShoppingCart struct {
//	Recipes []model.Recipe `json:"recipes"`
//}
//
//var shoppingCart ShoppingCart
//
//func deleteRecipeByID(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id, err := strconv.Atoi(params["id"])
//	if err != nil {
//		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
//		return
//	}
//
//	for index, recipe := range shoppingCart.Recipes {
//		if recipe.ID == id {
//			// Entferne das Rezept aus dem Warenkorb
//			shoppingCart.Recipes = append(shoppingCart.Recipes[:index], shoppingCart.Recipes[index+1:]...)
//			json.NewEncoder(w).Encode(recipe)
//			return
//		}
//	}
//
//	http.Error(w, "Recipe not found in shopping cart", http.StatusNotFound)
//}
//
//func SetupRoutes() {
//	router := mux.NewRouter()
//
//	router.HandleFunc("/shoppingcart/{id}", deleteRecipeByID).Methods("DELETE")
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}
