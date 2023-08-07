package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteShoppingCart(c *gin.Context) {
	if _, err := h.shoppingCardInMemoryStorage.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Löschen des Warenkorbs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warenkorb erfolgreich geleert"})
}
