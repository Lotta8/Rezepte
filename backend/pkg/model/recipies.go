package model

type Recipies struct { // Bauplan definiert
	Id   int
	Task string
	Done bool
}

func New(task string) Recipies {
	return Recipies{Task: task}
}

func (r *Recipies) Toggle() {
	r.Done = !r.Done
}
