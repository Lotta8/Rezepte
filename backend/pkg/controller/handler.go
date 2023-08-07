package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
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
	api.DELETE("/:id", h.DeleteShoppingCart)
	api.DELETE("/shoppingcart/:id", h.DeleteRecipeFromCart)
}

var ErrRecipeNotFound = errors.New("Rezept nicht im Warenkorb gefunden")

func (h *Handler) Run() {
	err := h.engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
