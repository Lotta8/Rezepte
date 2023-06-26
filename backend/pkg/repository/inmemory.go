package repository

import (
	"errors"
	"reminders/pkg/model"
)

type InMemoryStorage struct {
	Rezepte map[int]model.Recipe
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Rezepte: make(map[int]model.Recipe),
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
