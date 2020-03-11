package exemplos

import "fmt"

func variaveis() {
	// Definir variavel sem o var -> idade_new := 27

	var (
		nome   = "Victor"
		idade  = 22
		altura = 167
	)

	fmt.Println("O meu nome é", nome, "e tenho", idade, "anos")
	fmt.Println("Minha altura é", altura)

}
