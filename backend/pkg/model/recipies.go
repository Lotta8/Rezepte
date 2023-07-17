package model

type Ingredient struct {
	ID          int    `json:"id"`
	Bezeichnung string `json:"bezeichnung"`
}

type Recipe struct {
	ID          int                `json:"id"`
	Bezeichnung string             `json:"name"`
	Zutaten     []RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	Zutat   Ingredient `json:"zutat"`
	Einheit string     `json:"einheit"`
	Menge   float64    `json:"menge"`
}
