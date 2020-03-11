package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo. Erro: ", err.Error)
		return
	}

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler arquivo com o leitor CSV: ", err)
	}

	arquivoJSON, err := os.Create()
	if err != nil {
		fmt.Println("Houve um erro ao criar o arquivo JSON. Erro: ", err.Error)
		return
	}

	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s \n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s \n", indiceItem, item)
		}
	}

	arquivo.Close()
}
