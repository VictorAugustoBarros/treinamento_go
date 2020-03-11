package exemplos

import "fmt"

func whileloop() {
	var soma = 2

	for soma < 600 { // while
		soma += soma
	}

	fmt.Println(soma)
}
