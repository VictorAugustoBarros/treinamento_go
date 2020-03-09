package main

import (
	"encoding/json"
	"fmt"
	"udemy/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25

	if err := casa.SetValor(100); err != nil {
		fmt.Println("[main] Houve um erro na atribuição do valor a casa: ", err)
		if err == model.ErrValorDeImoveMuitoAlto {
			fmt.Println("Corretor, por favor verifique sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v \n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error)
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
