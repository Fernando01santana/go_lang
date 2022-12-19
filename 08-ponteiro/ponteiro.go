package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	variavel++

	fmt.Println(variavel, variavel2)

	//PONTEIRO E UMA REFERENCIA DE MEMORIA
	var variavel3 int = 0
	var ponteiro *int //informa que e um ponteiro do tipo int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3 //aponta para o endereco de memoria da variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //desfazendo o apontamento do endereco de memoria

}
