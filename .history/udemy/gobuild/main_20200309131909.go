package main

import (
	"encoding/json"
	"fmt"
	"udemy/mapas/model"
)

/*
GOOS=windows GOARH=386 go build -o meuapp.exe
GOOS=linux GOARH=arm go build -o meuappraspberry -v
*/

func main() {
	teste := model.Imovel
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100)

	fmt.Printf("Casa é: %+v \n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
