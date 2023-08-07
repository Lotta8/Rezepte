package controller

import (
	"errors"
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

	if _, err := h.shoppingCardInMemoryStorage.DeleteById(id); err != nil {
		if errors.Is(err, ErrRecipeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Rezept nicht im Warenkorb gefunden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Löschen des Rezepts aus dem Warenkorb"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rezept erfolgreich aus dem Warenkorb entfernt"})
}
