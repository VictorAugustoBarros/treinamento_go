package exemplos

import "fmt"

func switchExemple() {
	var nome = "João"

	switch nome {
	case "Ana":
		fmt.Println("É a Ana")
	case "João":
		fmt.Println("É o João")
	default:
		fmt.Println("Não conheço")
	}

	var nota = 4
	switch true {
	case nota > 9:
		fmt.Println("Ótimo")
	case nota > 7:
		fmt.Println("Muito bem")
	case nota > 6:
		fmt.Println("Bom")
	default:
		fmt.Println("Péssimo")
	}

}
