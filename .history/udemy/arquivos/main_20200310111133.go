package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	arquivo, err := os.Open("cidades.csv")
	if err != nil{
		fmt.Println("Houve um erro ao abrir o arquivo. Erro: ", err.Error)
		return
	}

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan(){
		linha := scanner.Text()
		fmt.Println(linha)
	}

	arquivo.Close()
}