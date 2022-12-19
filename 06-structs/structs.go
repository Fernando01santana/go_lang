package main

import "fmt"

// criando um tipo
type usuario struct {
	nome  string
	idade int8
}

func main() {
	var usuarioStruct usuario
	usuarioStruct.idade = 21
	usuarioStruct.nome = "Fernando"

	usuario2Struct := usuario{"Fernando Santana", 21}

	fmt.Println(usuario2Struct)
	fmt.Println(usuarioStruct)
}
