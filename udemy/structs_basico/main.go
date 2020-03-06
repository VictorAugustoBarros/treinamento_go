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
	fmt.Printf("O apartamento é: %+v \n", apartamento)

	chacara := Imovel{
		Y: 85,		
		Nome: "Chacara",		
		X: 22,
		valor: 55,		
	}
	fmt.Printf("A chacara é: %+v \n", chacara)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v \n", casa)

}