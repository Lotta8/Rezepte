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
func (w *Warenkorb) RezeptLöschen(rezeptID int) {
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

func main() {
	// Beispiel-Daten
	rezept1 := &Rezept{ID: 1, Name: "Würziger Fleischsalat"}
	rezept2 := &Rezept{ID: 2, Name: "Griessklösschen"}
	rezept3 := &Rezept{ID: 3, Name: "Spargelcremesuppe"}

	// Warenkorb erstellen und Rezepte hinzufügen
	warenkorb := &Warenkorb{
		Rezepte: []*Rezept{rezept1, rezept2, rezept3},
	}

	// Rezept mit ID 2 aus dem Warenkorb löschen
	warenkorb.RezeptLöschen(2)

	// Überprüfen, ob das Rezept gelöscht wurde
	for _, rezept := range warenkorb.Rezepte {
		fmt.Println(rezept.Name)
	}
}
