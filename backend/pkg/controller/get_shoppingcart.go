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
