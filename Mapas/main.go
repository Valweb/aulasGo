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

	apto := model.Imovel{}
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache?", len(cache)) // o len ler o comprimento de strings, mapas, verificar itens no cache, dados...

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	_, achou = cache["Casa Amarela"] // O underline só testa se encontrou o item no array ou não
	if achou {
		delete(cache, "Casa Amarela")
	}
	fmt.Println("Quantos itens há no cache?", len(cache))

}
