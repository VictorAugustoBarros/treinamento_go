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
