package repository

import (
	"errors"
)

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
	Menge   int
}

type InMemoryStorage struct {
	Zutaten map[int]Ingredient
	Rezepte map[int]Recipe
	nextID  int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Zutaten: make(map[int]Ingredient),
		Rezepte: make(map[int]Recipe),
		nextID:  1,
	}
}

func (s *InMemoryStorage) AddZutat(zutat Ingredient) {
	s.Zutaten[zutat.ID] = zutat
}

func (s *InMemoryStorage) GetZutat(id int) (Ingredient, error) {
	zutat, ok := s.Zutaten[id]
	if !ok {
		return Ingredient{}, errors.New("Zutat nicht gefunden")
	}
	return zutat, nil
}

func (s *InMemoryStorage) AddRezept(rezept Recipe) {
	rezept.ID = s.nextID
	s.nextID++
	s.Rezepte[rezept.ID] = rezept
}

func (s *InMemoryStorage) GetRezept(id int) (Recipe, error) {
	rezept, ok := s.Rezepte[id]
	if !ok {
		return Recipe{}, errors.New("Rezept nicht gefunden")
	}
	return rezept, nil
}

func (s *InMemoryStorage) AddZutatToRezept(rezeptID, zutatID int, einheit string, menge int) error {
	rezept, err := s.GetRezept(rezeptID)
	if err != nil {
		return err
	}

	zutat, err := s.GetZutat(zutatID)
	if err != nil {
		return err
	}

	recipeIngredient := RecipeIngredient{
		Zutat:   zutat,
		Einheit: einheit,
		Menge:   menge,
	}

	rezept.Zutaten = append(rezept.Zutaten, recipeIngredient)

	return nil
}

/*package repository

import (
	"fmt"
	"reminders/pkg/model"
)

var nextId = 2

type InMemoryStorageImpl struct {
	reminders *[]model.Reminder
}

func New() *InMemoryStorageImpl {
	nextId = 1
	reminders := make([]model.Reminder, 0)
	return &InMemoryStorageImpl{reminders: &reminders}
}

func (storage InMemoryStorageImpl) Get(id int) (*model.Reminder, error) {
	for _, reminder := range *storage.reminders {
		if id == reminder.Id {
			return &reminder, nil
		}
	}
	return nil, fmt.Errorf("id not exist")
}

func (storage InMemoryStorageImpl) GetAll() []model.Reminder {
	return *storage.reminders
}

func (storage InMemoryStorageImpl) Delete(id int) bool {
	for i, reminder := range *storage.reminders {
		if id == reminder.Id {
			*storage.reminders = append((*storage.reminders)[:i], (*storage.reminders)[i+1:]...)
			return true
		}
	}
	return false
}

func (storage InMemoryStorageImpl) Create(task string) (int, error) {
	reminder := model.New(task)
	reminder.Id = nextId
	nextId++
	*storage.reminders = append(*storage.reminders, reminder)
	return reminder.Id, nil
}

func (storage InMemoryStorageImpl) Update(id int) (*model.Reminder, error) {
	for i, reminder := range *storage.reminders {
		if id == reminder.Id {
			(&(*storage.reminders)[i]).Toggle()
			return &(*storage.reminders)[i], nil
		}
	}
	return nil, fmt.Errorf("id not exist")
}*/
