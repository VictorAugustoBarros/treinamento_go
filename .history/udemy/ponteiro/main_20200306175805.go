package main

import "fmt"

//Imovel -> struct que armazena os dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("Casa é: %p -  %+v \n", &casa, casa)
}
