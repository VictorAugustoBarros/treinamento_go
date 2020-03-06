package model

/*
Imovel -> representa informações de um imóvel
*/
type Imovel struct {
	X 		int
	Y 		int
	Nome 	string
	valor 	int
}

/*
SetValor -> defini o valor do imível
*/
func (i *Imovel) SetValor(valor int){
	i.valor = valor
}