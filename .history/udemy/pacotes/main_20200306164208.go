package main

import (
	"alura/pacotes/operadora"
	"alura/pacotes/prefixo"
	"fmt"
)

var (
	//NomeDoUsuario é nome do usuário do sistema
	NomeDoUsuario = "Victor"
)

func main() {
	fmt.Printf("Nome do usuário %s \n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d \n", prefixo.Capital)
	fmt.Printf("Prefixo da Capital: %s \n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s \n", prefixo.TesteComPrefixo)
}
