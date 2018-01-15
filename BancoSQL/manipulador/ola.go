package manipulador

import  (
"github.com/Valweb/aulasGo/ServidorWeb/model"
"net/http"
"time"
"fmt"
)
//Ola é o manipulador da requisição a rota /ola
func Ola(w http.ResponseWriter, r *http.Request){
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModelosOla.ExecuteTemplate(w, "ola.html", pagina); err != nil{
		http.Error(w, "Erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execução do modelo: ", err.Error())
	}
}

