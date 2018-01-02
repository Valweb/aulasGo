package main

import "fmt"

func main() {
	capitais := []string{"Lisboa"}
	//fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	//fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tokio"
	cidades[4] = "Singapura"
	//fmt.Println(cidades, len(cidades), cap(cidades))

	capitais[1] = "Brasilia"
	// o que tem no numeros, qual o comprimento, qual a capacidade
	fmt.Println(capitais, len(capitais), cap(capitais))
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}
	//Primeiro item começa com indice 0
	//Segundo item começa com o indice 1
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia, len(capitaisAsia))
	//um slice de outro slice
	//vazio significa início
	temp1 := cidades[:2]
	fmt.Println(temp1)
	//Comprimento tiro os dois do final
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)

	//Retirar um indice
	indiceDoItemAretirar := 2
	//Criar um slice temporário com o conteudo sem o item
	temp := cidades[:indiceDoItemAretirar]
	//Depois adicionar em cidades
	temp = append(temp, cidades[indiceDoItemAretirar+2:]...)
	copy(cidades, temp)
	fmt.Println(cidades)

}
