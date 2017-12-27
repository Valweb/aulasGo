package main

import "fmt"

func main() {
	var numeros []int
	// o que tem no numeros, qual o comprimento, qual a capacidade - Mas o slice não está inicializado
	fmt.Println(numeros, len(numeros), cap(numeros))

	// make inicializa o slice
	numeros = make([]int, 5)
	fmt.Println(numeros, len(numeros), cap(numeros))

	//O que tem, tamanho, capacidade
	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	//append para adicionar
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Paris"
	cidades[3] = "Tokio"
	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {
		fmt.Printf("cidade[%d] = %s\n\r", indice, cidade)
	}
}
