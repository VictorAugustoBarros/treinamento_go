package main

import (
	"encoding/json"
	"fmt"
	"udemy/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("Casa é: %+v \n", casa)
	fmt.Printf("O valor da casa é: %d \n", casa.GetValor())
	objJson, _ := json.Marshal(casa)
	fmt.Printf(objJson)
}
