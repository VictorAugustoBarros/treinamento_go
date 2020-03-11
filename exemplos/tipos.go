package exemplos

import "fmt"

func tipos() {
	var altura float64 = 1.67
	var aberto = true

	fmt.Printf("Tipo: %T Valor: %v\n", altura, altura)
	fmt.Printf("Tipo: %T Valor: %v\n", aberto, aberto)
}
