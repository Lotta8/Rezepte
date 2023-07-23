package controller

import (
	"fmt"
)

type Rezept struct {
	ID   int
	Name string
	// Weitere Rezeptdetails
}

type Warenkorb struct {
	Rezepte []*Rezept
}

// Funktion zum Löschen eines Rezepts aus dem Warenkorb
func (w *Warenkorb) RezeptLoeschen(rezeptID int) {
	for i, rezept := range w.Rezepte {
		if rezept.ID == rezeptID {
			// Rezept aus dem Warenkorb entfernen
			w.Rezepte = append(w.Rezepte[:i], w.Rezepte[i+1:]...)
			fmt.Println("Rezept erfolgreich gelöscht")
			return
		}
	}
	fmt.Println("Rezept nicht gefunden")
}
