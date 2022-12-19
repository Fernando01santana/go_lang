package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
}

type estudante struct {
	dados     pessoa
	matricula int8
}

func main() {
	pessoa1 := pessoa{"Fernando Santana", 21}
	pessoa2 := estudante{pessoa1, 18}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
