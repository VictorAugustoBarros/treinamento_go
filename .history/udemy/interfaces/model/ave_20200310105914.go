package model

/*
Galinha -> Representa uma ave do tipo Galinha
*/
type Galinha interface {
	Cacareja() string
}

/*
Pata -> Representa uma ave do tipo Pato
*/
type Pata interface {
	Quack() string
}

/*
Ave -> Representa um animal
*/
type Ave struct {
	Nome string
}

/*
Cacareja -> Retorna o som emitido por uma galinha
*/
func (a Ave) Cacareja() string {
	return "Cócóricó"
}

/*
Quack -> Retorna o som emitido por uma pata
*/
func (a Ave) Quack() string {
	return "quack, quack..."
}
