package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Valweb/aulasGo/EscreverArquivos/model"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	//Logo depois que abrir um arquivo colocar o defer, comando para fechar o arquivo
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
		return
	}

	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo. Erro: ", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			// strings.Split do pacote strings quebra o item
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v,\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)

			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item", item, " Erro:", err.Error())
				return
			}
			escritor.WriteString(" " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
	//Não precisa desses fechamentos porque estou usando o defer na linha 22 para fechar o arquivo
	//arquivoJSON.Close()
	//arquivo.Close()
}
