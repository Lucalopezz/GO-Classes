package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	fmt.Println("Enviando dados...")
	ch <- 42 // Envia dados ao canal
	fmt.Println("Dados enviados.")
}

func receiveData(ch chan int) {
	time.Sleep(500 * time.Millisecond) // Simula algum tempo de processamento
	val := <-ch                        // Recebe dados do canal
	fmt.Println("Recebido:", val)
}

func main() {
	// A comunicação entre goroutines é feita por canais.
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go sendData(ch) // Inicia uma goroutine para enviar dados
	}

	time.Sleep(1 * time.Second) // Espera um segundo para garantir que todos os envios terminem

	for i := 0; i < 5; i++ {
		go receiveData(ch) // Inicia uma goroutine para receber dados
	}

	time.Sleep(1 * time.Second) // Espera um segundo para garantir que todos os recebimentos terminem
}
