package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker iniciado: ", id)
	time.Sleep(time.Second)
	fmt.Println("Worker finalizado: ", id)
}

func workerWithErrorGroup(id int) error {
	fmt.Println("Worker iniciado: ", id)
	if id == 2 {
		return errors.New("Erro no worker: " + fmt.Sprint(id))
	}
	time.Sleep(time.Second)
	fmt.Println("Worker finalizado: ", id)

	return nil
}

func main() {
	fmt.Println("--------------------------------Wait Group ------------------------")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Fim")

	fmt.Println("--------------------------------Error Group ------------------------")

	// Error Group é uma forma de gerenciar goroutines e capturar erros de forma mais elegante do que usando WaitGroup diretamente.
	// Não precisa de passar um WaitGroup, pois o errgroup já gerencia isso internamente.
	// nem do contador de goroutines
	var eg errgroup.Group

	eg.SetLimit(3) // Limita o número de goroutines simultâneas a 3
	// Ele espera que uma go routine termine para começar a próxima, caso o limite seja atingido.

	for i := 1; i <= 3; i++ {
		eg.Go(func() error {
			return workerWithErrorGroup(i)
		})
	}

	err := eg.Wait()
	// Se alguma das goroutines retornar um erro, o Wait() retornará esse erro.
	// se mais de uma goroutine retornar um erro, o Wait() retornará o primeiro erro encontrado.
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Fim")
	}
}
