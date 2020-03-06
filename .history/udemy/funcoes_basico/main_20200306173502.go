package main

import (
	"fmt"
	"udemy/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d \n", resultado)

	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d \n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
}
