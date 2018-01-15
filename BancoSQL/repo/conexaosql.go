package repo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mysql"   //baixar esse pacote

)
// DB armazena a conexão com o banco
var Db *sqlx.DB

//Funcão que abre conexão com Mysql
func AbreConexaoBancoDeDadosSQL() (err error){

	err = nill
	Db, err = sqlx.Open("mysql", "root@cp(127.0.0.1:3306/cursogo")
		if err != nil{
			return
		}
		err = Db.Ping(){

			if err != nil{
				return
			}
		}


}