package main

import (
	"bufio"
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
		return
	}

	arquivoJSON, err := os.Create()
	if err != nil {
		fmt.Println("Houve um erro ao criar o arquivo JSON. Erro: ", err.Error)
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {		
		for indiceItem, item := range linha {
			dados := 
		}
	}

	arquivo.Close()
}
