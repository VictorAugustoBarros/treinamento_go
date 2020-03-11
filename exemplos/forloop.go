package exemplos

import "fmt"

func forloop() {
	var soma int = 0

	for i := 1; i <= 10; i++ {
		soma = soma + i
	}

	fmt.Println(soma)
}
