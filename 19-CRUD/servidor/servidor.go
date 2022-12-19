package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type book struct {
	name string `json:"name"`
	id   string `json:"id"`
}

// criar usuario - insere
func CriarLivro(response http.ResponseWriter, request *http.Request) {

	//recebe as informacoes contidas no body da requisicao
	corpoRequisicao, erro := ioutil.ReadAll(request.Body)

	//verifica se foi enviada alguma informacao
	if erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Certifique-se de que enviou todas as informaçoes corretamente e tente novamente."))
		return
	}

	//verifica se a conversao para json foi bem sucedida
	var book book
	if erro = json.Unmarshal(corpoRequisicao, &book); erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao transformar usuario para struct"))
		return
	}

	//verifica se foi possivel conectar com o banco
	db, erro := banco.Conectar()
	if erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	//protege contra mysqlinjector criando o comando de inserçao sem inserir os valores
	statement, erro := db.Prepare("insert into books (name) values (?)")

	//verifica se conseguiu criar o statement
	if erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao criar statement"))
		return
	}

	//fecha a conexao com o banco e executa o statement
	defer statement.Close()
	insercao, erro := statement.Exec(book.name)

	if erro != nil {
		fmt.Println(erro)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao criar statement"))
		return
	}

	idBookInsert, erro := insercao.LastInsertId()
	if erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao recuperar book inserido"))
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(fmt.Sprintf("id do usuario inserido: %d", idBookInsert)))
	return

}

// buscar todos os livros - list
func BuscarLivros(response http.ResponseWriter, request *http.Request) {
	//verifica se foi possivel conectar com o banco
	db, erro := banco.Conectar()
	if erro != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from books")
	if erro != nil {
		response.Write([]byte("Erro ao recuperar books"))
		return
	}
	defer linhas.Close()

	var books []book

	for linhas.Next() {
		var book book
		if erro := linhas.Scan(&book.id, &book.name); erro != nil {
			response.Write([]byte("Erro ao escanear book"))
			return
		}
		books = append(books, book)
	}

	response.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(response).Encode(books); erro != nil {
		response.Write([]byte("Erro ao converter books para json"))
		return
	}
}

// buscar um unico livro - select
func BuscarLivro(response http.ResponseWriter, request *http.Request) {

}
