package repository

import (
	"errors"
	"recipies/pkg/model"
)

func New() *InMemoryStorage {
	// Implementierung der New-Funktion, um ein neues InMemoryStorage-Objekt zu erstellen
	return &InMemoryStorage{}
}

var WürzigerFleischsalat = model.Recipe{
	ID:          1,
	Bezeichnung: "Würziger Fleischsalat",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          1,
				Bezeichnung: "Fleischwurst",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          2,
				Bezeichnung: "Schinkenwurst",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          3,
				Bezeichnung: "Essiggurken",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          4,
				Bezeichnung: "Mayonnaise",
			},
			Einheit: "g",
			Menge:   150,
		},
		{
			Zutat: model.Ingredient{
				ID:          5,
				Bezeichnung: "Miracle Whip",
			},
			Einheit: "g",
			Menge:   80,
		},
		{
			Zutat: model.Ingredient{
				ID:          6,
				Bezeichnung: "Senf",
			},
			Einheit: "TL",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          7,
				Bezeichnung: "Pfeffer",
			},
			Einheit: "Pr",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "TL",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   1,
		},
	},
}

var Griesskloesschen = model.Recipe{
	ID:          2,
	Bezeichnung: "Griessklösschen",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "EL",
			Menge:   2,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   250,
		},
		{
			Zutat: model.Ingredient{
				ID:          12,
				Bezeichnung: "Hartweizengriess",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          13,
				Bezeichnung: "Petersilie",
			},
			Einheit: "Pck",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          14,
				Bezeichnung: "Eier",
			},
			Einheit: "Stk",
			Menge:   2,
		},
	},
}
var Spargelcremesuppe = model.Recipe{
	ID:          3,
	Bezeichnung: "Spargelcremesuppe",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          15,
				Bezeichnung: "weisse Spargeln",
			},
			Einheit: "kg",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          16,
				Bezeichnung: "Schalotten",
			},
			Einheit: "Stk",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "Pr",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "EL",
			Menge:   3/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          18,
				Bezeichnung: "Rahm",
			},
			Einheit: "dl",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "EL",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          19,
				Bezeichnung: "Gemüsebouillon",
			},
			Einheit: "dl",
			Menge:   1/2,
		},
	},
}
var GulaschMitSpätzle = model.Recipe{
	ID:          4,
	Bezeichnung: "Gulasch mit Spätzle",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          20,
				Bezeichnung: "Gulasch",
			},
			Einheit: "g",
			Menge:   200,
		},
		{
			Zutat: model.Ingredient{
				ID:          21,
				Bezeichnung: "Zwiebel",
			},
			Einheit: "Stk",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          22,
				Bezeichnung: "Butterschmalz",
			},
			Einheit: "EL",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          23,
				Bezeichnung: "Tomatenmark",
			},
			Einheit: "EL",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          24,
				Bezeichnung: "Majoran",
			},
			Einheit: "EL",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          25,
				Bezeichnung: "Wacholderbeeren",
			},
			Einheit: "TL",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          26,
				Bezeichnung: "Wein, rot",
			},
			Einheit: "ml",
			Menge:   75,
		},
		{
			Zutat: model.Ingredient{
				ID:          27,
				Bezeichnung: "Fleischbrühe",
			},
			Einheit: "ml",
			Menge:  125,
		},
		{
			Zutat: model.Ingredient{
				ID:          28,
				Bezeichnung: "Honig",
			},
			Einheit: "TL",
			Menge:  1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          29,
				Bezeichnung: "Spätzle",
			},
			Einheit: "g",
			Menge:  100,
		},
	},
}
var KartoffelgratinMitPilzragout = model.Recipe{
	ID:          5,
	Bezeichnung: "Kartoffelgratin mit Pilzragout",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          30,
				Bezeichnung: "Kartoffeln",
			},
			Einheit: "g",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          18,
				Bezeichnung: "Rahm",
			},
			Einheit: "ml",
			Menge:   100,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   25,
		},
		{
			Zutat: model.Ingredient{
				ID:          31,
				Bezeichnung: "Pilze",
			},
			Einheit: "g",
			Menge:   50,
		},
		{
			Zutat: model.Ingredient{
				ID:          16,
				Bezeichnung: "Schalotten",
			},
			Einheit: "g",
			Menge:   12,
		},
		{
			Zutat: model.Ingredient{
				ID:          32,
				Bezeichnung: "Greyerzer",
			},
			Einheit: "g",
			Menge:   25,
		},
		{
			Zutat: model.Ingredient{
				ID:          33,
				Bezeichnung: "Knoblauchzehe",
			},
			Einheit: "Stk",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:  5,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:  1,
		},
		{
			Zutat: model.Ingredient{
				ID:          7,
				Bezeichnung: "Pfeffer",
			},
			Einheit: "Pr",
			Menge:  1,
		},
		{
			Zutat: model.Ingredient{
				ID:          34,
				Bezeichnung: "Muskat",
			},
			Einheit: "Pr",
			Menge:  1,
		},
	},
}
var MilchReisMitKirschen = model.Recipe{
	ID:          6,
	Bezeichnung: "Milchreis mit Kirschen",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          35,
				Bezeichnung: "Milchreis",
			},
			Einheit: "g",
			Menge:   75,
		},
		{
			Zutat: model.Ingredient{
				ID:          36,
				Bezeichnung: "Kirschen",
			},
			Einheit: "g",
			Menge:   88,
		},
		{
			Zutat: model.Ingredient{
				ID:          37,
				Bezeichnung: "Zimt",
			},
			Einheit: "Pr",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          38,
				Bezeichnung: "Speisestärke",
			},
			Einheit: "EL",
			Menge:   1/2,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   380,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "EL",
			Menge:   3/4,
		},
	},
}
var Apfelstrudel = model.Recipe{
	ID:          7,
	Bezeichnung: "Apfelstrudel",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          39,
				Bezeichnung: "Äpfel",
			},
			Einheit: "g",
			Menge:   225,
		},
		{
			Zutat: model.Ingredient{
				ID:          40,
				Bezeichnung: "Zitrone",
			},
			Einheit: "Stk",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          41,
				Bezeichnung: "Sultaninen",
			},
			Einheit: "g",
			Menge:   13,
		},
		{
			Zutat: model.Ingredient{
				ID:          42,
				Bezeichnung: "Rohzucker",
			},
			Einheit: "g",
			Menge:   23,
		},
		{
			Zutat: model.Ingredient{
				ID:          43,
				Bezeichnung: "Vanillezucker",
			},
			Einheit: "Pck",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          37,
				Bezeichnung: "Zimt",
			},
			Einheit: "TL",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   25,
		},
		{
			Zutat: model.Ingredient{
				ID:          12,
				Bezeichnung: "Hartweizengriess",
			},
			Einheit: "EL",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "dl",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   75,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "TL",
			Menge:   1/4,
		},
	},
}
var Kaiserschmarrn = model.Recipe{
	ID:          8,
	Bezeichnung: "Kaiserschmarrn",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   38,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "dl",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "EL",
			Menge:   3/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          44,
				Bezeichnung: "Eigelb",
			},
			Einheit: "Stk",
			Menge:   3/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          45,
				Bezeichnung: "Eiweiss, steif",
			},
			Einheit: "Stk",
			Menge:   3/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   13,
		},
		{
			Zutat: model.Ingredient{
				ID:          46,
				Bezeichnung: "Rosinen",
			},
			Einheit: "g",
			Menge:   10,
		},
	},
}
var Fasnachtskrapfen = model.Recipe{
	ID:          9,
	Bezeichnung: "Fasnachtskrapfen",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   6,
		},
		{
			Zutat: model.Ingredient{
				ID:          14,
				Bezeichnung: "Eier",
			},
			Einheit: "Stk",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          44,
				Bezeichnung: "Eigelb",
			},
			Einheit: "Stk",
			Menge:   1/4,
		},
		{
			Zutat: model.Ingredient{
				ID:          47,
				Bezeichnung: "Konfitüre",
			},
			Einheit: "g",
			Menge:   25,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   50,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   25,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   1,
		},
		{
			Zutat: model.Ingredient{
				ID:          48,
				Bezeichnung: "Puderzucker",
			},
			Einheit: "g",
			Menge:   5,
		},
	},
}

fmt.Printf("Default-Rezept: %+v\n", defaultRecipe)
}


type InMemoryStorage struct {
	Rezepte []model.Recipe
	nextID  int
}

func NewInMemoryStorage() *InMemoryStorage {
	rezepte := make([]model.Recipe, 0)
	rezepte = append(rezepte, WürzigerFleischsalat, Griessklösschen, Spargelcremesuppe, GulaschMitSpätzle, KartoffelgratinMitPilzragout, MilchReisMitKirschen, Apfelstrudel, Kaiserschmarrn, Fasnachtskrapfen)
	return &InMemoryStorage{
		Rezepte: rezepte,
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


func (s *InMemoryStorage) AddZutatToRezept(rezeptID, zutatID int, einheit string, menge float64) error {
	_, err := s.GetRezept(rezeptID)
	if err != nil {
		return err
	}
	return nil
}
