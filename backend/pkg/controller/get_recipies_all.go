package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecipeOverviewResponse struct {
	ID          int    `json:"id"`
	Bezeichnung string `json:"Bezeichnung"`
}

func (h *Handler) GetRecipes(c *gin.Context) {
	recipes := h.recipeInMemoryStorage.GetAllRecipe()

	recipeOverviewResponse := make([]RecipeOverviewResponse, 0)
	for _, v := range recipes {
		recipeOverviewResponse = append(recipeOverviewResponse, RecipeOverviewResponse{ID: v.ID, Bezeichnung: v.Bezeichnung})
	}

	c.JSON(http.StatusOK, recipeOverviewResponse)
}
