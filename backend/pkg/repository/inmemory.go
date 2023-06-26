package repository

import (
	"errors"
	"fmt"
	"reminders/pkg/model"
)

var suppe = model.Recipe{
	ID: 1,
	Bezeichnung: "Würziger Fleischsalat",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.in{
				ID:          1,
				Bezeichnung: "Zutat 1",
			},
			Einheit: "g",
			Menge: 10
				,
		},
		{
			Zutat: model.Ingredient{
				ID:          2,
				Bezeichnung: "Zutat 2",
			},
			Einheit: "ml",
			Menge:   250,
		},
	},
}

var Salat = model.Recipe{
	ID: 1,
	Bezeichnung: "Würziger Fleischsalat",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          3,
				Bezeichnung: "Zutat 3",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          2,
				Bezeichnung: "Zutat 2",
			},
			Einheit: "ml",
			Menge:   250,
		},
	},
}

// Ausgabe der Default-Variable
fmt.Printf("Default-Rezept: %+v\n", defaultRecipe)
}







}
		//"Schinkenwurst", "Essiggurken", "Mayonnaise", "Miracle Whip", " Senf", "Pfeffer", "Salz", "Zucker"}},
type InMemoryStorage struct {
	Rezepte []model.Recipe
}

func NewInMemoryStorage() *InMemoryStorage {
	rezepte := make([]model.Recipe, 0)
	rezepte = append(rezepte, suppe, Salat)
	return &InMemoryStorage{
		Rezepte: rezepte
	}
}

func (s *InMemoryStorage) AddRezept(rezept model.Recipe) {
	rezept.ID = s.nextID
	s.nextID++
	s.Rezepte[rezept.ID] = rezept
}

func (s *InMemoryStorage) GetRezept(id int) (model.Recipe, error) {
	rezept, ok := s.Rezepte[id]
	if !ok {
		return model.Recipe{}, errors.New("Rezept nicht gefunden")
	}
	return rezept, nil
}

func (s *InMemoryStorage) AddZutatToRezept(rezeptID, zutatID int, einheit string, menge int) error {
	rezept, err := s.GetRezept(rezeptID)
	if err != nil {
		return err
	}
}
