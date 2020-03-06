package main

import (
	"fmt"
	"udemy/structs_avancado/model"
)

func main() {
	casa := model.Imovel
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.setValor(60000)

	fmt.Printf("Casa é: %+v \n", casa)
}
