package main



import (
	"fmt"
	"net/http"

	"github.com/Valweb/aulasGo/BancoSQL/manipulador"
	"github.com/Valweb/aulasGo/BancoSQL/repo"
)

func init() {
	//Essa mensagem vai iniciar antes do main, pq está na função init
	fmt.Println("Vamos começar a subir o servidor...")  
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	//Rotas
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}