package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	cliente := &http.Client{
		//Quando criar um cliente http é importante definir o seu timeout
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Google Brasil. Erro:", err.Error())
		return
	}
	//defer para quando acaber encerrar tudo
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		//Usar o ioutil pq a página vai retorar texto(html)
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google Brasil. Erro:", err.Error())
			return
		}

		fmt.Println(corpo)
	}

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o Google Brasil. Erro:", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste") //São chaves criptografadas que o dono da API passa
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Google Brasil. Erro:", err.Error())
		return
	}
	//defer para quando acaber encerrar tudo
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		//Usar o ioutil pq a página vai retorar texto(html)
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google Brasil. Erro:", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(corpo)
	}

}
