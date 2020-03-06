package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "Piraquara"

	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Piraquara" {
		fmt.Println("O cliente vive em Piraquara")
	}

	if descricao, status := aQuantoTempoEstaDevendo(meses);

}

func aQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}
	descricao = "O cliente esta com o carnê em dia."
	return
}
