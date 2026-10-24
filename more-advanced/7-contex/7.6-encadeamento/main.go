package main

import (
	"context"
	"fmt"
	"time"
)

type myString string

const (
	timestamp myString = "timestamp"
	id        myString = "id"
)

func escreverContexto(ctx context.Context, id string) {
	fmt.Println("Registro de tempo:", ctx.Value(timestamp))
	fmt.Println("ID:", ctx.Value(id))
}

func main() {
	ctxPai, cancelPai := context.WithCancel(
		context.WithValue(context.Background(), timestamp, time.Now().Format(time.RFC3339)),
	)

	ctxFilho, _ := context.WithCancel(
		context.WithValue(ctxPai, id, "123"),
	)

	ctxFilho2, cancelFilho2 := context.WithCancel(
		context.WithValue(ctxFilho, id, "456"),
	)

	fmt.Println("Contexto pai:")
	escreverContexto(ctxPai, "id não existe no contexto pai")

	fmt.Println("Contexto filho:")
	escreverContexto(ctxFilho, "id do contexto filho")

	fmt.Println("Contexto filho 2:")
	escreverContexto(ctxFilho2, "id do contexto filho 2")

	go func() {
		<-ctxPai.Done()
		fmt.Println("Contexto pai cancelado")
	}()

	go func() {
		<-ctxFilho.Done()
		fmt.Println("Contexto filho1 cancelado")
	}()

	go func() {
		<-ctxFilho2.Done()
		fmt.Println("Contexto filho2 cancelado")
	}()

	time.Sleep(2 * time.Second)
	cancelFilho2()
	fmt.Println("Cancelando contexto filho2")

	// quando o pai encerra, todos os filhos também são encerrados
	time.Sleep(2 * time.Second)
	cancelPai()
	fmt.Println("Cancelando contexto pai")

	time.Sleep(2 * time.Second)
}
