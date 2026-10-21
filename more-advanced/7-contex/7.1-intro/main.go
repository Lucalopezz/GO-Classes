package main

import (
	"context"
	"fmt"
)

type (
	mySring1 string
	mySring2 string
)

func main() {
	// Context é uma forma de passar informaçãoes compartilhadas
	// controlar o ciclo de vide de operações em um programa
	ctx := context.Background()

	// para por um valor dentro do contexto é necessário criar um novo contexto com o valor desejado, utilizando a função WithValue.
	ctx = context.WithValue(ctx, mySring1("key"), "value")

	fmt.Println(ctx.Value("key"))      // value
	fmt.Println(ctx.Value("otherKey")) // <nil>

	fmt.Println("chamando a função leitura")
	leitura(ctx)
}

// no go é convencional que funções que recebem um contexto como parâmetro, o primeiro parâmetro seja o contexto.

func leitura(ctx context.Context) {
	if v := ctx.Value("key"); v != nil {
		fmt.Println(v)
		return
	}
	fmt.Println("key não encontrada")
}
