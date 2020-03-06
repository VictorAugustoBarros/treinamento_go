package main

import "udemy/structs_avancado/model"

func main() {
	casa := model.Imovel
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.setValor()
}
