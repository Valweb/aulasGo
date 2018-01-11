package main

import (
	"fmt"
	"net/http"
	

	"github.com/Valweb/aulasGo/ServidorWeb/manipulador"
)

func main() {
	//Criar um servidor Web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo!")
				
	})
	

	//Rotas
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/css/", fs)

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando, comandante... =D")
	http.ListenAndServe(":8181", nil,) //Escute e sirva
}
