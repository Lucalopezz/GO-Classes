package main

import (
	"fmt"
	"time"
)

// Só é possível enviar dados ao canal chan<-
func sendData(ch chan<- int) {
	fmt.Println("Enviando dados...")
	ch <- 42
	fmt.Println("Dados enviados.")
}

// Só é possível receber dados do canal <-chan
func receiveData(ch <-chan int) {
	time.Sleep(500 * time.Millisecond)
	val := <-ch
	fmt.Println("Recebido:", val)
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go sendData(ch)
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 5; i++ {
		go receiveData(ch) // Inicia uma goroutine para receber dados
	}

	time.Sleep(1 * time.Second) // Espera um segundo para garantir que todos os recebimentos terminem
}
