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
	Quack()
}

type Ave struct {
	Nome string
}

func (a Ave) Cacareja() string {
	return "Cócóricó"
}
