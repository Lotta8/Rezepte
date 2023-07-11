package repository

import (
	"errors"
	"recipies/pkg/model"
	"sync"
)

type InMemoryStorage struct {
	mutex        sync.RWMutex
	Rezepte      map[int]model.Recipe
	nextID       int
	cartByUserID map[int][]string
}

func NewInMemoryStorage() *InMemoryStorage {
	rezepte := make(map[int]model.Recipe)
	return &InMemoryStorage{
		Rezepte:      rezepte,
		nextID:       1,
		cartByUserID: make(map[int][]string),
	}
}

func (s *InMemoryStorage) AddRezept(rezept model.Recipe) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rezept.ID = s.nextID
	s.nextID++
	s.Rezepte[rezept.ID] = rezept
}

func (s *InMemoryStorage) GetRezept(id int) (model.Recipe, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	rezept, ok := s.Rezepte[id]
	if !ok {
		return model.Recipe{}, errors.New("Rezept nicht gefunden")
	}
	return rezept, nil
}

func (s *InMemoryStorage) AddZutatToRezept(rezeptID, zutatID int, einheit string, menge float64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rezept, err := s.GetRezept(rezeptID)
	if err != nil {
		return err
	}

	zutat := model.Ingredient{ID: zutatID}
	recipeIngredient := model.RecipeIngredient{Zutat: zutat, Einheit: einheit, Menge: menge}
	rezept.Zutaten = append(rezept.Zutaten, recipeIngredient)
	s.Rezepte[rezeptID] = rezept

	return nil
}

func (s *InMemoryStorage) DeleteRecipe(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.Rezepte[id]
	if !ok {
		return errors.New("Rezept nicht gefunden")
	}

	delete(s.Rezepte, id)

	return nil
}

func (s *InMemoryStorage) DeleteEinkaufskorb(userID int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.cartByUserID[userID]
	if !ok {
		return errors.New("Einkaufskorb nicht gefunden")
	}

	delete(s.cartByUserID, userID)

	return nil
}

func (s *InMemoryStorage) AddToCart(userID int, item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cartByUserID[userID] = append(s.cartByUserID[userID], item)
}

func (s *InMemoryStorage) GetCart(userID int) ([]string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cart, ok := s.cartByUserID[userID]
	if !ok {
		return nil, errors.New("Einkaufswagen nicht gefunden")
	}

	return cart, nil
}

func (s *InMemoryStorage) RemoveFromCart(userID int, item string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, ok := s.cartByUserID[userID]
	if !ok {
		return errors.New("Einkaufswagen nicht gefunden")
	}

	for i, v := range cart {
		if v == item {
			// Remove the item from the cart
			s.cartByUserID[userID] = append(cart[:i], cart[i+1:]...)
			return nil
		}
	}

	return errors.New("Artikel nicht im Einkaufswagen gefunden")
}

var WuerzigerFleischsalat = model.Recipe{
	ID:          1,
	Bezeichnung: "Würziger Fleischsalat",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          1,
				Bezeichnung: "Fleischwurst",
			},
			Einheit: "g",
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          2,
				Bezeichnung: "Schinkenwurst",
			},
			Einheit: "g",
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          3,
				Bezeichnung: "Essiggurken",
			},
			Einheit: "g",
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          4,
				Bezeichnung: "Mayonnaise",
			},
			Einheit: "g",
			Menge:   150.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          5,
				Bezeichnung: "Miracle Whip",
			},
			Einheit: "g",
			Menge:   80.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          6,
				Bezeichnung: "Senf",
			},
			Einheit: "TL",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          7,
				Bezeichnung: "Pfeffer",
			},
			Einheit: "Pr",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "TL",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   1.00,
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
			Menge:   2.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   250.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          12,
				Bezeichnung: "Hartweizengriess",
			},
			Einheit: "g",
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          13,
				Bezeichnung: "Petersilie",
			},
			Einheit: "Pck",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          14,
				Bezeichnung: "Eier",
			},
			Einheit: "Stk",
			Menge:   2.00,
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
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          16,
				Bezeichnung: "Schalotten",
			},
			Einheit: "Stk",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "Pr",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "EL",
			Menge:   0.75,
		},
		{
			Zutat: model.Ingredient{
				ID:          18,
				Bezeichnung: "Rahm",
			},
			Einheit: "dl",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "EL",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          19,
				Bezeichnung: "Gemüsebouillon",
			},
			Einheit: "dl",
			Menge:   0.50,
		},
	},
}

var GulaschMitSpaetzle = model.Recipe{
	ID:          4,
	Bezeichnung: "Gulasch mit Spätzle",
	Zutaten: []model.RecipeIngredient{
		{
			Zutat: model.Ingredient{
				ID:          20,
				Bezeichnung: "Gulasch",
			},
			Einheit: "g",
			Menge:   200.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          21,
				Bezeichnung: "Zwiebel",
			},
			Einheit: "Stk",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          22,
				Bezeichnung: "Butterschmalz",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          23,
				Bezeichnung: "Tomatenmark",
			},
			Einheit: "EL",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          24,
				Bezeichnung: "Majoran",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          25,
				Bezeichnung: "Wacholderbeeren",
			},
			Einheit: "TL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          26,
				Bezeichnung: "Wein, rot",
			},
			Einheit: "ml",
			Menge:   75.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          27,
				Bezeichnung: "Fleischbrühe",
			},
			Einheit: "ml",
			Menge:   125.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          28,
				Bezeichnung: "Honig",
			},
			Einheit: "TL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          29,
				Bezeichnung: "Spätzle",
			},
			Einheit: "g",
			Menge:   100.00,
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
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          18,
				Bezeichnung: "Rahm",
			},
			Einheit: "ml",
			Menge:   100.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   25.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          31,
				Bezeichnung: "Pilze",
			},
			Einheit: "g",
			Menge:   50.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          16,
				Bezeichnung: "Schalotten",
			},
			Einheit: "g",
			Menge:   12.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          32,
				Bezeichnung: "Greyerzer",
			},
			Einheit: "g",
			Menge:   25.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          33,
				Bezeichnung: "Knoblauchzehe",
			},
			Einheit: "Stk",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   5.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          7,
				Bezeichnung: "Pfeffer",
			},
			Einheit: "Pr",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          34,
				Bezeichnung: "Muskat",
			},
			Einheit: "Pr",
			Menge:   1.00,
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
			Menge:   75.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          36,
				Bezeichnung: "Kirschen",
			},
			Einheit: "g",
			Menge:   88.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          37,
				Bezeichnung: "Zimt",
			},
			Einheit: "Pr",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          38,
				Bezeichnung: "Speisestärke",
			},
			Einheit: "EL",
			Menge:   0.50,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "ml",
			Menge:   380.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "EL",
			Menge:   0.75,
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
			Menge:   225.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          40,
				Bezeichnung: "Zitrone",
			},
			Einheit: "Stk",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          41,
				Bezeichnung: "Sultaninen",
			},
			Einheit: "g",
			Menge:   13.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          42,
				Bezeichnung: "Rohzucker",
			},
			Einheit: "g",
			Menge:   23.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          43,
				Bezeichnung: "Vanillezucker",
			},
			Einheit: "Pck",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          37,
				Bezeichnung: "Zimt",
			},
			Einheit: "TL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   25.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          12,
				Bezeichnung: "Hartweizengriess",
			},
			Einheit: "EL",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "dl",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   75.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "TL",
			Menge:   0.25,
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
			Menge:   38.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          11,
				Bezeichnung: "Milch",
			},
			Einheit: "dl",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          9,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          8,
				Bezeichnung: "Zucker",
			},
			Einheit: "EL",
			Menge:   0.75,
		},
		{
			Zutat: model.Ingredient{
				ID:          44,
				Bezeichnung: "Eigelb",
			},
			Einheit: "Stk",
			Menge:   0.75,
		},
		{
			Zutat: model.Ingredient{
				ID:          45,
				Bezeichnung: "Eiweiss, steif",
			},
			Einheit: "Stk",
			Menge:   0.75,
		},
		{
			Zutat: model.Ingredient{
				ID:          10,
				Bezeichnung: "Butter",
			},
			Einheit: "g",
			Menge:   13.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          46,
				Bezeichnung: "Rosinen",
			},
			Einheit: "g",
			Menge:   10.00,
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
			Menge:   6.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          14,
				Bezeichnung: "Eier",
			},
			Einheit: "Stk",
			Menge:   1.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          17,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   60.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          47,
				Bezeichnung: "Mandeln",
			},
			Einheit: "g",
			Menge:   4.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          48,
				Bezeichnung: "Zucker",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          49,
				Bezeichnung: "Salz",
			},
			Einheit: "Pr",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          50,
				Bezeichnung: "Rum",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          51,
				Bezeichnung: "Mehl",
			},
			Einheit: "g",
			Menge:   50.00,
		},
		{
			Zutat: model.Ingredient{
				ID:          52,
				Bezeichnung: "Puderzucker",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          53,
				Bezeichnung: "Zitronensaft",
			},
			Einheit: "EL",
			Menge:   0.25,
		},
		{
			Zutat: model.Ingredient{
				ID:          54,
				Bezeichnung: "Zimt",
			},
			Einheit: "Pr",
			Menge:   1.00,
		},
	},
}

var storage *InMemoryStorage

func InitializeStorage() *InMemoryStorage {
	if storage == nil {
		storage = NewInMemoryStorage()
		storage.AddRezept(WuerzigerFleischsalat)
		storage.AddRezept(Griesskloesschen)
		storage.AddRezept(Spargelcremesuppe)
		storage.AddRezept(GulaschMitSpaetzle)
		storage.AddRezept(KartoffelgratinMitPilzragout)
		storage.AddRezept(MilchReisMitKirschen)
		storage.AddRezept(Apfelstrudel)
		storage.AddRezept(Kaiserschmarrn)
		storage.AddRezept(Fasnachtskrapfen)
	}
	return storage
}
