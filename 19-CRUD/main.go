package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", servidor.CriarLivro).Methods(http.MethodPost)
	router.HandleFunc("/books/list", servidor.BuscarLivros).Methods(http.MethodGet)
	fmt.Println("Executando na porta 5002")
	log.Fatal(http.ListenAndServe(":5002", router))
}
