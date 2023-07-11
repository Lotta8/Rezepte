package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"recipies/pkg/repository"
	"strconv"
)

type Handler struct {
	engine     *gin.Engine
	repository *repository.InMemoryStorage
}

func NewHandler() *Handler {
	h := &Handler{
		engine:     gin.Default(),
		repository: repository.NewInMemoryStorage(),
	}

	h.setupRoutes()

	return h
}

func (h *Handler) setupRoutes() {
	api := h.engine.Group("/api/recipe")

	api.GET("/:id", h.GetRecipe)
	api.DELETE("/:id", h.DeleteRecipe)
	api.GET("/all", h.GetRecipes)
	api.POST("/", h.CreateRecipe)
	api.PUT("/", h.UpdateRecipe)
	api.DELETE("/shoppingcart/:id", h.DeleteEinkaufskorb)
}

func (h *Handler) Run() {
	err := h.engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	err = h.repository.DeleteRecipe(recipeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}

func (h *Handler) GetRecipe(c *gin.Context) {
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	recipe, err := h.repository.GetRecipe(recipeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

func (h *Handler) GetRecipes(c *gin.Context) {
	recipes := h.repository.GetAllRecipes()
	c.JSON(http.StatusOK, recipes)
}

func (h *Handler) CreateRecipe(c *gin.Context) {
	var recipe repository.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newRecipe := h.repository.CreateRecipe(&recipe)
	c.JSON(http.StatusCreated, newRecipe)
}

func (h *Handler) UpdateRecipe(c *gin.Context) {
	var recipe repository.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedRecipe, err := h.repository.UpdateRecipe(&recipe)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, updatedRecipe)
}

func (h *Handler) DeleteEinkaufskorb(c *gin.Context) {
	id := c.Param("id")
	recipeID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	err = h.repository.DeleteEinkaufskorb(recipeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Einkaufskorb not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Einkaufskorb deleted successfully"})
}
