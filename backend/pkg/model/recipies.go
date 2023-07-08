package model

type Ingredient struct {
	ID          int
	Bezeichnung string
}

type Recipe struct {
	ID          int
	Bezeichnung string
	Zutaten     []RecipeIngredient
}

type RecipeIngredient struct {
	Zutat   Ingredient
	Einheit string
	Menge   float64
}
