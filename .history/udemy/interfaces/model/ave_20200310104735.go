package model

/*
Galinha -> Representa uma ave do tipo Pato
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
