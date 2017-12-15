package main

import (
	"fmt"

	"github.com/Valweb/aulasGo/Structs_Avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0) // Inicilização do cache

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela bonita dos seus sonhos"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa Amarela no cash?")
	Imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, o que achei foi: %v\n\r", Imovel)
	}
}
