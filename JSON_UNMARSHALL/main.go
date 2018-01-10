package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Valweb/aulasGo/JSON_UNMARSHALL/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	//Servidor API para testes
	resposta, err := cliente.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Google Brasil. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do Google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o servidor. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Servidor. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do servidor. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON do servidor. Erro: ", err.Error())
			return
		}
		//montar os campos juntos na estrutura com %+v, e quebrar lina \r\n
		fmt.Printf("Post é: %+v\r\n", post)
	}
}
