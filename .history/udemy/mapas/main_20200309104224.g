package main

import "udemy/structs_avancado/model"

var cache map[string]model.Imovel

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)
}