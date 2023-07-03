package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"recipies/pkg/repository"
)

type Handler struct { // Objekt Handler
	engine     *gin.Engine                 // Datentyp gin.engine
	repository *repository.InMemoryStorage // selber definiert
}

func New() *Handler {
	h := &Handler{ // wir erstellen einen neuen Handler
		engine:     gin.Default(), // Eigenschaften auf Zeilen 10 + 11 definiert.
		repository: repository.New(),
	}
	router := h.engine.Group("/api/recipe/") // Gruppe definiert, in dieser Gruppe folgende Endpunkte definiert, damit ich die Service aufrufen kann. h ist das objekt, welches ich oben instanziert habe.
	router.GET("/:id", h.GetRecipe)          // Punkt ist immer von einem Objekt, welches instanziert wurde. damit kann z.B. auf die Eigenschaften, Function oder Services zugegriffen werden.
	router.DELETE("/:id", h.DeleteRecipe)    // Rauts definieren
	router.GET("/all", h.GetRecipes)
	router.POST("/", h.CreateRecipe)
	router.PUT("/", h.UpdateRecipe)
	router.DELETE("/", h.DeleteEinkaufskorb)
	return h
}

func (h *Handler) Run() { // Func welches an Objekt Handler gebunden ist.
	err := h.engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) DeleteRecipe(context *gin.Context) {

}

func (h *Handler) GetRecipe(context *gin.Context) {

}

func (h *Handler) GetRecipes(context *gin.Context) {

}
