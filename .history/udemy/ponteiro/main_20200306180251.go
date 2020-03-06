package main

import "fmt"

//Imovel -> struct que armazena os dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("Casa é: %p -  %+v \n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara Linda", 280000}
	fmt.Printf("Chacara é: %p -  %+v \n", &chacara, chacara)
}

/*
mudaImovel -> teste
*/
func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 1
}
