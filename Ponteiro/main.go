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
	casa := new(Imovel)
	fmt.Printf("Casa é %p - %v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara linda", 2784940}
	fmt.Printf("Chacara é %p - %v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Chacara é %p - %v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5

}
