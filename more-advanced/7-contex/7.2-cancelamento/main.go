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
			fmt.Println("Trabalhador encerrando", id)
			return
		default:
			fmt.Println("Trabalhador executando", id)
		}
	}
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelando trabalhadores")
	cancel()
	time.Sleep(1 * time.Second)
}
