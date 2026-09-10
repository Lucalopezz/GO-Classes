/*
Exercício 03 — Baixa de estoque

Implemente a rotina de venda de uma loja. O estoque é um map em que o código
do produto aponta para um produto armazenado por ponteiro.

A função deve:

1. Retornar erro se o código não existir.
2. Retornar erro se a quantidade solicitada for menor ou igual a zero.
3. Retornar erro se não houver estoque suficiente.
4. Diminuir o estoque e retornar o valor total da venda em centavos.

Crie também uma função para imprimir o estoque restante e faça testes com
vendas válidas e inválidas. Use struct, map, ponteiro, método ou função
auxiliar e tratamento explícito de erros.
*/
package main

type Product struct {
	Code       string
	Name       string
	Stock      int
	PriceCents int
}

func processSale(inventory map[string]*Product, code string, quantity int) (int, error) {
	// TODO: implemente o exercício.
	panic("not implemented")
}

func ex03() {
	// TODO: monte um estoque, processe vendas e exiba os erros quando ocorrerem.
}
