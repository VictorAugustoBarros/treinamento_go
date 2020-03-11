package exemplos

import "fmt"

func coercao() {
	var a int = 46
	var b float64 = 6.4

	var c int = int(b)

	fmt.Printf("Tipo: %T Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T Valor: %v\n", b, c)
	fmt.Printf("Tipo: %T Valor: %v\n", c, c)
}
