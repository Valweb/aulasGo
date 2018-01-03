package model

//Galinha representa uma ave
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja Retorna o som de uma galinha
func (a Ave) Cacareja() string {
	return "Cocoricó..."
}

// Quack Retorna o som emitido por um pato
func (a Ave) Quack() string {
	return "Quack, Quack..."
}
