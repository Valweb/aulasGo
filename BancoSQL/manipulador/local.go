package manipulador

import (

	"fmt"
	"net/http"
	"strconv"
	"github.com/Valweb/aulasGo/BancoSQL/model"
	"github.com/Valweb/aulasGo/BancoSQL/repo"
)


//Local é o manipulador da requisição de rota /local/
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
		codigoTelefone, err := strconv.Atoi(r.URL.Path[7:]) 

	if err != nil {
		http.Error(w, "Não foi enviado um número válido!", http.StatusBadRequest)
		fmt.Println("[Local]Erro ao converter o numero enviado ", err.Error())
		return
	}

	//Query
	sql := "select country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Query(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse numero", http.StatusInternalServerError)
		fmt.Println("[Local]não foi possível executar a pesquisa ", sql, "Erro:"err.Error())
		return
	}

	for linha.Next(){
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar esse numero", http.StatusInternalServerError)
			fmt.Println("[Local]não foi possível fazer o binding dosdo banco na struct  ",err.Error())
			return
		}

	}


	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}
}

