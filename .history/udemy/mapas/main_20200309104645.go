package main

import (
	"fmt"
	"udemy/mapas/model"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa
	fmt.Println("Há uma casa amarela no cache?")
}
