package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteRecipeFromCart(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Rezept-ID"})
		return
	}

	success, err := h.shoppingCardInMemoryStorage.DeleteById(id)
	if err != nil || !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rezept nicht im Warenkorb gefunden"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rezept erfolgreich aus dem Warenkorb entfernt"})
}
