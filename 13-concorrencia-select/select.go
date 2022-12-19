package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)
	//executa antes de qualquer coisa
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	//executa antes de qualquer coisa
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//senario usado apenas com concorrencia
		select {
		case mensagemcanal1 := <-canal1:
			fmt.Println(mensagemcanal1)

		case mensagemcanal2 := <-canal2:
			fmt.Println(mensagemcanal2)
		}
	}
}
