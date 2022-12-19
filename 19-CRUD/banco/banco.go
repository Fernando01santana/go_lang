package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	connection := "root:12345@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", connection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	fmt.Println("Conexao esta aberta!")
	return db, nil

}
