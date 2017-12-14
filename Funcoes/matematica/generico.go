package matematica

//Calculo executa qualquer tipo de calculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x * y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetua a divisão
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto que retorna o resto de uma divisão
func DivisorComResto(numeroA, int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
