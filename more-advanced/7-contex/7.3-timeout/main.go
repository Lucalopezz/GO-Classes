package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker encerrando", id)
			return
		default:
			fmt.Println("Worker executando", id)
		}
	}
}

func main() {
	ctx := context.Background()

	// quando o tempo limite é atingido, o contexto é cancelado e os workers param de executar
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// quando a data limite é atingida, o contexto é cancelado e os workers param de executar
	// ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(4 * time.Second)
}
