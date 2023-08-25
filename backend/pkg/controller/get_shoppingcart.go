package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShoppingCartItemsResponse struct {
	Recipes     []RecipeOverviewResponse `json:"recipes"`
	Ingredients []IngredientResponse     `json:"ingredients"`
}

type IngredientResponse struct {
	Menge       float64 `json:"menge"`
	Bezeichnung string  `json:"bezeichnung"`
	Einheit     string  `json:"einheit"`
}

func (h *Handler) GetShoppingCart(context *gin.Context) {
	shoppingCartItems := h.shoppingCardInMemoryStorage.Get()

	ingredientMap := make(map[string]float64)
	recipesMap := make(map[int]RecipeOverviewResponse)
	ingredientUnits := make(map[string]string)

	for _, item := range shoppingCartItems {
		for _, ingredient := range item.Recipe.Zutaten {
			ingredientName := ingredient.Zutat.Bezeichnung
			if _, exists := ingredientUnits[ingredientName]; !exists {
				ingredientUnits[ingredientName] = ingredient.Einheit
			}
			ingredientMap[ingredientName] += ingredient.Menge
		}
		if _, exists := recipesMap[item.Recipe.ID]; !exists {
			recipesMap[item.Recipe.ID] = RecipeOverviewResponse{
				ID:          item.Recipe.ID,
				Bezeichnung: item.Recipe.Bezeichnung,
			}
		}
	}

	var ingredients []IngredientResponse
	for ingredientName, quantity := range ingredientMap {
		ingredients = append(ingredients, IngredientResponse{
			Menge:       quantity,
			Bezeichnung: ingredientName,
			Einheit:     ingredientUnits[ingredientName],
		})
	}

	var recipes []RecipeOverviewResponse
	for _, recipe := range recipesMap {
		recipes = append(recipes, recipe)
	}

	shoppingCartResponse := ShoppingCartItemsResponse{
		Recipes:     recipes,
		Ingredients: ingredients,
	}

	context.JSON(http.StatusOK, shoppingCartResponse)
}
