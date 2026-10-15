package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	dados []int
	cond  *sync.Cond
	mu    sync.Mutex
}

func main() {
	buffer := Buffer{
		dados: make([]int, 0),
		// nao inicializa a condicao de sincronizacao aqui, pois precisamos passar um locker (mutex) para ela
	}

	// Recebe um locker (interface com um Locker e um Unlocker) e cria uma nova condição de sincronização associada a esse locker.
	buffer.cond = sync.NewCond(&buffer.mu)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			buffer.mu.Lock()
			fmt.Println("Produzindo:", i)
			buffer.dados = append(buffer.dados, i)

			buffer.cond.Signal() // Notifica a condição de sincronização que um novo item foi adicionado ao buffer

			buffer.mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			buffer.mu.Lock()

			// for para evitar o problema de "spurious wakeup" (acordar sem motivo) e garantir que o buffer não esteja vazio antes de consumir um item
			for len(buffer.dados) == 0 {
				buffer.cond.Wait() // Aguarda a condição de sincronização ser sinalizada
			}

			item := buffer.dados[0]
			buffer.dados = buffer.dados[1:] // Remove o item do buffer -> recria uma nova slice sem o item consumido
			fmt.Println("Consumindo:", item)

			buffer.mu.Unlock()
		}
	}()

	time.Sleep(2 * time.Second)
}
