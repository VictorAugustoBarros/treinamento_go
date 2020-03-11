package main

import (
	"fmt"
	"udemy/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

}

func QueroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja)
}
