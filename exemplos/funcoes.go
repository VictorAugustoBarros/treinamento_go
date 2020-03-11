package exemplos

func soma(x int, y int) int {
	return x + y
}

// func calcular(a int) (quadrado int, cubo int) {
// 	quadrado = a * a
// 	cubo = a * a * a

// 	return
// }

func calcular(a int) (int, int) {
	var quadrado = a * a
	var cubo = a * a * a

	return quadrado, cubo
}
