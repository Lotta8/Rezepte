package model

type Recipies struct { // Bauplan definiert
	Ingredien Ingredien
}

type Ingredien struct {
	Id          int
	Bezeichnung string
	Menge       int
	Unit        Unit
}

type Unit struct {
	Id          int
	Bezeichnung string
}

func New(task string) Recipies {
	return Recipies{Task: task}
}

func (r *Recipies) Toggle() {
	r.Done = !r.Done
}
