package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecipeRequest struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

func (h *Handler) AddRecipeToCart(context *gin.Context) {
	var recipeRequest RecipeRequest

	err := context.BindJSON(&recipeRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Request data not valid"})
		return
	}

	_, err = h.shoppingCardInMemoryStorage.Add(recipeRequest.ID, recipeRequest.Count)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	context.Status(http.StatusOK)
}
