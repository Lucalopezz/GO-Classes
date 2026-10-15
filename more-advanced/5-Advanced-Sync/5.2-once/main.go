package main

import (
	"fmt"
	"sync"
	"time"
)

type BancoDeDados struct {
	conectado bool
}

var (
	bancoDeDados *BancoDeDados
	// Tem que ser uma variável global para que todas as goroutines compartilhem a mesma instância de sync.Once.
	once sync.Once
)

func ObterBancoDeDados() *BancoDeDados {
	// Varias goroutines podem chamar ObterBancoDeDados simultaneamente, mas a conexão será estabelecida apenas uma vez.
	once.Do(func() {
		fmt.Println("Conectando ao banco de dados...")
		time.Sleep(2 * time.Second) // Simula o tempo de conexão
		bancoDeDados = &BancoDeDados{conectado: true}
		fmt.Println("Conexão estabelecida!")
	})

	return bancoDeDados
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d solicitando conexão ao banco de dados...\n", id)
			db := ObterBancoDeDados()
			if db.conectado {
				fmt.Printf("Goroutine %d recebeu a conexão ao banco de dados!\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Todas as goroutines concluídas.")
}
