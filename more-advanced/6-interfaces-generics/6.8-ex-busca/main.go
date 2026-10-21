package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Pessoa struct {
	Nome  string
	Idade int
}

type Buscaable interface {
	constraints.Integer | constraints.Float | Pessoa | ~string
}

// Busca
// Função que busque um item em uma lista de int, string ou Pessoa, e retorne o índice do item encontrado ou -1 se não encontrado.

func Busca[T Buscaable](lista []T, item T) int {
	if len(lista) == 0 {
		return -1
	}
	for i, v := range lista {
		if v == item {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Busca de inteiros:", Busca([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println("Busca de strings:", Busca([]string{"banana", "maçã", "laranja"}, "maçã"))
	fmt.Println("Busca de pessoas:", Busca([]Pessoa{
		{Nome: "João", Idade: 30},
		{Nome: "Maria", Idade: 25},
	}, Pessoa{Nome: "Maria", Idade: 25}))
}
