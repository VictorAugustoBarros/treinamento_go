package matematica

/*
Calculo -> executa qualquer tipo de calculo, basta enviar a funcao desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

/*
Multiplicador -> multiplica x vezes y
*/
func Multiplicador(x int, y int) int {
	return x * y
}

/*
Divisor -> divide x vezes y
*/
func Divisor(x int, y int) (resultado int) {
	resultado = x / y
	return
}