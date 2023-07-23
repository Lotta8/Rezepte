package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"recipies/pkg/model"
	"recipies/pkg/repository"
)

type Handler struct {
	engine                      *gin.Engine
	recipeInMemoryStorage       *repository.RecipeInMemoryStorage
	shoppingCardInMemoryStorage *repository.ShoppingCartInMemoryStorage
}

func NewHandler() *Handler {
	h := &Handler{
		engine:                      gin.Default(),
		recipeInMemoryStorage:       repository.NewRecipeInMemoryStorage(),
		shoppingCardInMemoryStorage: repository.NewShoppingCartInMemoryStorage(),
	}

	h.setupRoutes()

	return h
}

func (h *Handler) setupRoutes() {
	api := h.engine.Group("/api")
	api.GET("/recipe/:id", h.GetRecipe)
	api.GET("/recipe/all", h.GetRecipes)

	api.POST("/shopping-cart", h.AddRecipeToCart)
	api.GET("/shopping-cart", h.GetShoppingCart)
	// api.DELETE("/:id", h.DeleteRecipe)
	// api.DELETE("/shoppingcart/:id", h.DeleteEinkaufskorb)
}

func (h *Handler) Run() {
	err := h.engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

//func (h *Handler) DeleteRecipe(c *gin.Context) {
//	id := c.Param("id")
//	recipeID, err := strconv.Atoi(id)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
//		return
//	}
//
//	err = h.recipeInMemoryStorage.DeleteRecipe(recipeID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
//}

// Ihr habt keine UpdateRecipe Function im Repository.
// Frage: Braucht ihr das wirklich für eure Projekt? Die Rezepte sind bei euch fix gegeben und ihr müsst diese doch nur lesen?
// GetRezeptDetail(id int) / GetRezpet(id int)
// GetAllRezept()
func (h *Handler) UpdateRecipe(c *gin.Context) {
	var recipe model.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

}

//func (h *Handler) DeleteEinkaufskorb(c *gin.Context) {
//	id := c.Param("id")
//	recipeID, err := strconv.Atoi(id)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
//		return
//	}
//
//	err = h.recipeInMemoryStorage.DeleteEinkaufskorb(recipeID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Einkaufskorb not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Einkaufskorb deleted successfully"})
//}
