package controller

import (
	"errors"
	"github.com/gin-contrib/cors"
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
	configureCors(h.engine)
	api := h.engine.Group("/api")
	api.GET("/recipe/:id", h.GetRecipe)
	api.GET("/recipe/all", h.GetRecipes)

	api.POST("/shopping-cart", h.AddRecipeToCart)
	api.GET("/shopping-cart", h.GetShoppingCart)
	api.DELETE("/shopping-cart", h.DeleteShoppingCart)
	api.DELETE("/shopping-cart/:id", h.DeleteRecipeFromCart)
}

var ErrRecipeNotFound = errors.New("Rezept nicht im Warenkorb gefunden")

func (h *Handler) Run() {
	err := h.engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func configureCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	engine.Use(cors.New(config))
}
