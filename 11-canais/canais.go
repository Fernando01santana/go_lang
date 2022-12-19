package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	//nao espera funcao para poder seguir
	go escrever("Olá Mundo!", canal)
	fmt.Println("Depois da funcao escrever comecar a ser executada")

	//recebendo valor do canal e salvando na variavel mensagem
	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//enviando valor para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}

//deadlock = e quando o canal fica esperando receber um valor e nao recebe
