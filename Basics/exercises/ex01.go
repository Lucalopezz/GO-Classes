/*
Exercício 01 — Relatório de despesas

Crie um relatório para uma lista de despesas do mês.

 1. Some o valor total gasto em cada categoria (alimentação, transporte,
    moradia etc.) e devolva esses totais em um map[string]float64.
 2. Encontre a despesa de maior valor.
 3. Retorne um erro caso exista uma despesa com valor negativo.

Use uma struct para representar cada despesa, uma função com múltiplos
retornos, slices, maps e um loop. No final, imprima um resumo organizado
por categoria e a maior despesa.
*/
package main

type Expense struct {
	Description string
	Category    string
	Amount      float64
}

func summarizeExpenses(expenses []Expense) (map[string]float64, Expense, error) {
	// TODO: implemente o exercício.
	panic("not implemented")
}

func ex01() {
	// TODO: crie dados de teste e chame summarizeExpenses.
}
