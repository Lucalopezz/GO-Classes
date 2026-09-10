/*
Exercício 05 — Fechamento de carrinho com políticas de desconto

Implemente o fechamento de um carrinho de compras usando centavos inteiros
para evitar problemas de precisão com float64.

 1. Some o subtotal de todos os itens.
 2. Valide que preço e quantidade não sejam negativos e que a quantidade seja
    maior que zero.
 3. Aplique uma política de desconto recebida por uma interface.
 4. Nunca permita que o total final fique negativo.

Crie pelo menos duas políticas concretas, por exemplo uma de percentual e
outra de cupom com valor fixo, implementando o método exigido pela interface.
Faça o checkout com cada política e também sem desconto. Retorne erros para
um carrinho inválido e use métodos sempre que eles deixarem o código mais
claro.
*/
package main

type CartItem struct {
	Description    string
	UnitPriceCents int
	Quantity       int
}

type DiscountPolicy interface {
	Discount(totalCents int) int
}

func checkout(items []CartItem, policy DiscountPolicy) (int, error) {
	// TODO: implemente o exercício.
	panic("not implemented")
}

func ex05() {
	// TODO: crie itens, implemente políticas de desconto e teste o checkout.
}
