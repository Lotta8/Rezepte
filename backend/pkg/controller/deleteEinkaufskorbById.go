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

//func main() {
//
//	rezept1 := &Rezept{ID: 1, Name: "Würziger Fleischsalat"}
//	rezept2 := &Rezept{ID: 2, Name: "Griessklösschen"}
//	rezept3 := &Rezept{ID: 3, Name: "Spargelcremesuppe"}
//	rezept4 := &Rezept{ID: 4, Name: "Gulasch mit Spätzle"}
//	rezept5 := &Rezept{ID: 5, Name: "Kartoffelgratin mit Pilzragout"}
//	rezept6 := &Rezept{ID: 6, Name: "Milchreis mit Kirschen"}
//	rezept7 := &Rezept{ID: 7, Name: "Apfelstrudel"}
//	rezept8 := &Rezept{ID: 8, Name: "Kaiserschmarrn"}
//	rezept9 := &Rezept{ID: 9, Name: "Fasnachtskrapfen"}
//
//	// Warenkorb erstellen und Rezepte hinzufügen
//	warenkorb := &Warenkorb{
//		Rezepte: []*Rezept{rezept1, rezept2, rezept3, rezept4, rezept5, rezept6, rezept7, rezept8, rezept9},
//	}
//
//	// Rezept mit ID 2 aus dem Warenkorb löschen
//	warenkorb.RezeptLoeschen(2)
//
//	// Überprüfen, ob das Rezept gelöscht wurde
//	for _, rezept := range warenkorb.Rezepte {
//		fmt.Println(rezept.Name)
//	}
//}
