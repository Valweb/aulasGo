package main

import (
	"encoding/json"
	"fmt"

	"github.com/Valweb/aulasGo/Structs_Avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(80000)

	fmt.Printf("Casa é: %v\r\n", casa)
	// ou da forma abaixo
	fmt.Printf("Casa é: %v\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	//Printf objJSON imprimeo json do arquivo imovel.go
	fmt.Printf("A casa em JSON:", string(objJSON))
}
