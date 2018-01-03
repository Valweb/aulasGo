package main

import (
	"fmt"

	"github.com/Valweb/aulasGo/Interface/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"
	QueroAcordarComCacarejo(jojo)
	QueroOuvirUmaPata(jojo)
}

func QueroAcordarComCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func QueroOuvirUmaPata(p model.Pata) {
	fmt.Println(p.Quack())
}
