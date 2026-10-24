package main

import (
	"fmt"
	"time"
)

// O padrão de broadcast é um padrão de concorrência em que uma única fonte de dados envia informações para múltiplos receptores simultaneamente.

func consumer(id int, ch <-chan int) {
	for val := range ch {
		fmt.Printf("Consumidor %d recebeu: %d\n", id, val)
	}
}

func broadcast(channels []chan int, value int) {
	for _, ch := range channels {
		ch <- value
	}
}

func main() {
	channals := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	for i, ch := range channals {
		go consumer(i+1, ch)
	}

	broadcast(channals, 42)

	time.Sleep(1 * time.Second) // Espera os consumidores terminarem o processamento
}
