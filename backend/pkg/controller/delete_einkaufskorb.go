package controller

import (
	"errors"
)

/*func (h Handler) DeleteReminder(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	deleted := h.repository.Delete(id)
	if deleted {
		context.Status(http.StatusOK)
	} else {
		context.JSON(http.StatusBadRequest, "could not delete reminder")
	}
}*

type EinkaufskorbService struct {
	storage repository.EinkaufskorbStorage
}

func NewEinkaufskorbService(storage repository.EinkaufskorbStorage) *EinkaufskorbService {
	return &EinkaufskorbService{
		storage: storage,
	}
}*/

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
