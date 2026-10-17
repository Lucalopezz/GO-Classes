package main

import (
	"bytes"
	"fmt"
	"sync"
)

var pool = sync.Pool{
	// So vai ser chamado quando o pool estiver vazio e for necessário criar um novo objeto
	New: func() any {
		fmt.Println("Creating a new buffer")
		// Retorna o novo objeto que será armazenado no poo (pointer para um bytes.Buffer)
		return new(bytes.Buffer)
	},
}

// O pool é uma estrutura de dados que armazena objetos que podem ser reutilizados, evitando a necessidade de alocação
// e desalocação frequente de memória. Ele é útil para melhorar o desempenho em situações onde a criação de objetos é cara ou frequente.
func main() {
	// faz o cast do objeto retornado pelo pool para o tipo correto (pointer para bytes.Buffer)
	buffer := pool.Get().(*bytes.Buffer)

	// adiciona uma string ao buffer
	buffer.WriteString("Hello, World!")

	fmt.Println(buffer.String())

	buffer.Reset()   // limpa o buffer antes de devolvê-lo ao pool
	pool.Put(buffer) // devolve o buffer para o pool

	fmt.Println("====== New buffer from pool ======")
	buffer2 := pool.Get().(*bytes.Buffer)
	buffer2.WriteString("Hello, Go!")

	fmt.Println(buffer2.String())
}
