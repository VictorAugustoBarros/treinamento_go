package main

import "fmt"

//Imovel -> struct que armazena os dados de um imovel
type Imovel struct{
	X 		int
	Y 		int 
	Nome	string
	valor 	int
}

func main(){
	casa := Imovel{}
	fmt.Printf("A casa é: %+v \n", casa)

	apartamento := Imovel{ 17, 56, "Meu Ap", 760000 }
	fmt.Printf("O apartamento é: %+v \n", casa)
}