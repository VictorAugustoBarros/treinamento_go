package exemplos

import "fmt"

func inferencia() {
	var i int
	var j = i

	var k = 0.6 + 0.3i

	fmt.Printf("Tipo: %T Valor: %v\n", i, i)
	fmt.Printf("Tipo: %T Valor: %v\n", j, j)
	fmt.Printf("Tipo: %T Valor: %v\n", k, k)
}
