package main

import "fmt"

//Imovel é um struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 5373839}
	fmt.Printf("O apartamento é: %v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}
	fmt.Printf("A chacara é: %v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.Y = 10
	casa.X = 33
	casa.valor = 200
	fmt.Printf("A casa é: %v\r\n", casa)
}
