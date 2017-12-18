package main

import (
	"encoding/json"
	"fmt"

	"github.com/Valweb/aulasGo/Erro/model2"
)

func main() {
	casa := model2.Imovel{}
	casa.Nome = "casa da Maria"
	casa.X = 17
	casa.Y = 35
	if err := casa.SetValor(100000000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor da casa:", err.Error())
		if err == model2.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique sua avaliação.")
		}

		return
	}

	fmt.Printf("Casa é %+v\r\n", casa)

	objJSON, err := json.Marshal(casa) // A função Marshal retorna um array de bytes e retorna um erro.

	if err != nil { //Nil = nulo
		fmt.Println("[main] Houve erro na geração do objeto erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON:", string(objJSON))

}
