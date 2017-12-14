package main

import (
	"fmt"

	"github.com/Valweb/aulasGo/Funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = somatoria(120, 320)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma do pacote matematica foi: %d\r\n", resultado)
	resultado = matematica.Divisor(12, 3)
	fmt.Printf("O resultado da divisão do pacote matematica foi: %d\r\n", resultado)
}

func somatoria(x int, y int) int {
	return x * y
}
