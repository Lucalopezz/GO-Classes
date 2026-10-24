package main

// sync.Map é uma implementação de map concorrente em Go que permite acesso seguro por várias goroutines sem bloqueio explícito.
// Ele faz parte do pacote sync e foi feito para situações em que várias goroutines precisam ler e escrever concorrentemente em um map compartilhado.

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map // Usa sync.Map para acesso concorrente ao map
		wg sync.WaitGroup
	)

	// sync.Map não é do mesmo tipo que um map comum; portanto, não podemos usar a mesma sintaxe para adicionar pares de chave e valor.
	// Em vez disso, usamos o método Store de sync.Map para adicionar um par de chave e valor ao map.

	// Espera as 100 goroutines terminarem antes de sair do programa
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			// Faz o mesmo que o mutex
			m.Store(i, i) // Usa o método Store de sync.Map para adicionar um par de chave e valor ao map
		}()
	}

	wg.Wait() // Espera todas as goroutines terminarem antes de sair do programa

	// for não funciona com sync.Map; por isso, usamos o método Range para percorrer o map
	m.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true // Continua percorrendo o map
	})

	value, ok := m.Load(50)
	if ok {
		fmt.Printf("Valor da chave 50: %d\n", value)
	} else {
		fmt.Println("Chave 50 não encontrada")
	}
}
