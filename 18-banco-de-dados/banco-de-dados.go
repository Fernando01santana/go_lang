package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConnect := "root:12345@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlConnect)
	defer db.Close()
	if erro != nil {
		log.Fatal(erro)
	}

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexao esta aberta!")
}
