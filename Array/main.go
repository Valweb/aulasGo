package main

import "fmt"

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("Qual a capacidade deste array?", len(teste))
	//Defino a capacidade, depois defino o tipo
	cantores := [2]string{"Michael Jackson", "Eric Clapton"}
	fmt.Println("O que há nesse Array? \n\r", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade desse Array", len(capitais))
	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\n\r", indice, cidade)
	}
}
