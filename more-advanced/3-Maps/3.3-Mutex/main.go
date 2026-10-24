package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	// Espera as 100 goroutines terminarem antes de sair do programa
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done() // Diminui o contador do WaitGroup quando a goroutine termina
			mu.Lock()       // Bloqueia o mutex para garantir acesso exclusivo ao map
			m[i] = i
			mu.Unlock() // Libera o mutex para permitir que outras goroutines acessem o map
			// Este Done diminui o contador do WaitGroup,
			// usado para sinalizar que uma goroutine terminou seu trabalho.
		}()
	}

	wg.Wait() // Espera todas as goroutines terminarem antes de sair do programa

	for k, v := range m {
		fmt.Printf("%d: %d\n", k, v)
	}
}
