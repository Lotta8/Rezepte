package controller

import (
	"errors"
	"fmt"
	"recipies/pkg/repository"
)

type EinkaufskorbService struct {
	storage *repository.ShoppingCart
}

type EinkaufskorbNotFoundError struct {
	UserID int
}

func (e *EinkaufskorbNotFoundError) Error() string {
	return fmt.Sprintf("Einkaufskorb für Benutzer %d nicht gefunden", e.UserID)
}

func NewEinkaufskorbService(storage *repository.ShoppingCart) *EinkaufskorbService {
	return &EinkaufskorbService{storage: storage}
}

func (s *EinkaufskorbService) DeleteEinkaufskorb(userID int) error {
	_, err := s.storage.GetEinkaufskorb(userID)
	if err != nil {
		return &EinkaufskorbNotFoundError{UserID: userID}
	}

	err = s.storage.DeleteEinkaufskorb(userID)
	if err != nil {
		return errors.New("Fehler beim Löschen des Einkaufskorbs")
	}

	return nil
}
