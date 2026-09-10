/*
Exercício 04 — Fila de atendimento prioritário

Crie uma função que retire da fila o próximo chamado a ser atendido.

As prioridades possíveis são "urgente", "alta" e "normal". A prioridade
mais alta deve ser atendida primeiro. Quando dois chamados tiverem a mesma
prioridade, preserve a ordem em que chegaram.

A função deve retornar o chamado escolhido, a fila restante e um erro para
uma fila vazia ou para uma prioridade inválida. Use switch para definir o
peso das prioridades e percorra o slice apenas uma vez para localizar o
próximo chamado.

Depois, simule vários atendimentos até a fila ficar vazia e mostre a ordem
em que os clientes foram chamados.
*/
package main

type Ticket struct {
	ID       int
	Client   string
	Priority string
}

func nextTicket(tickets []Ticket) (Ticket, []Ticket, error) {
	// TODO: implemente o exercício.
	panic("not implemented")
}

func ex04() {
	// TODO: crie uma fila de chamados e simule o atendimento.
}
