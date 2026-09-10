/*
Exercise 04 — Priority support queue

Create a function that removes the next ticket to be handled from the queue.

The possible priorities are "urgent", "high", and "normal". The highest
priority must be handled first. When two tickets have the same priority,
preserve the order in which they arrived.

The function must return the selected ticket, the remaining queue, and an
error for an empty queue or an invalid priority. Use switch to define the
priority weights and iterate over the slice only once to find the next ticket.

Then simulate multiple support interactions until the queue is empty and
display the order in which the customers were called.
*/
package main

type Ticket struct {
	ID       int
	Client   string
	Priority string
}

func nextTicket(tickets []Ticket) (Ticket, []Ticket, error) {
	// TODO: implement the exercise.
	panic("not implemented")
}

func ex04() {
	// TODO: create a ticket queue and simulate the support process.
}
