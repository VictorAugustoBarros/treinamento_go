package main

import "udemy/mapas/model"

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0))


	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)
}
