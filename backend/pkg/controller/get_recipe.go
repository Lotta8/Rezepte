package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetRecipe(c *gin.Context) {
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	recipe, err := h.recipeInMemoryStorage.GetRecipe(recipeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	// optional mapping von model zu response model...
	c.JSON(http.StatusOK, recipe)
}
