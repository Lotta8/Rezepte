package controller

import (
	"errors"
	"recipies/pkg/repository"
)

type EinkaufskorbService struct {
	storage *repository.ShoppingCart
}

func (s *EinkaufskorbService) DeleteEinkaufskorb(userID int) error {
	// Überprüfe, ob der Einkaufskorb für den Benutzer vorhanden ist
	_, err := s.storage.GetEinkaufskorb(userID)
	if err != nil {
		return errors.New("Einkaufskorb nicht gefunden")
	}

	// Lösche den Einkaufskorb
	err = s.storage.DeleteEinkaufskorb(userID)
	if err != nil {
		return errors.New("Fehler beim Löschen des Einkaufskorbs")
	}

	return nil
}
