package main

import (
	"fmt"
	"udemy/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

}

/*
QueroAcordarComUmCacarejo -> retorna o som da galinha
*/
func QueroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

/*
QueroOuvirUmaPataNoLago -> retorna o som da pata
*/
func QueroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
