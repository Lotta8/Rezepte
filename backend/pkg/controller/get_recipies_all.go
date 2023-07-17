package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecipeOverviewResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRecipes(c *gin.Context) {
	recipes := h.repository.GetAllRecipe()

	recipeOverviewResponse := make([]RecipeOverviewResponse, 0)
	for _, v := range recipes {
		recipeOverviewResponse = append(recipeOverviewResponse, RecipeOverviewResponse{ID: v.ID, Name: v.Bezeichnung})
	}

	c.JSON(http.StatusOK, recipeOverviewResponse)
}
