package main

//Imovel -> struct que armazena os dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmr.Printf("Casa é: %+v \n", casa)
}
