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
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v \n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul Amarela"
	apto.X = 18
	apto.Y = 25
	apto.SetValor(60000)

}
