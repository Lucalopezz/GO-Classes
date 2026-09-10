/*
Exercise 03 — Inventory deduction

Implement a store's sales routine. The inventory is a map where the product
code points to a product stored through a pointer.

The function must:

1. Return an error if the code does not exist.
2. Return an error if the requested quantity is less than or equal to zero.
3. Return an error if there is not enough stock.
4. Deduct the stock and return the sale's total value in cents.

Also create a function to print the remaining inventory and test both valid
and invalid sales. Use a struct, map, pointer, method or helper function,
and explicit error handling.
*/
package main

type Product struct {
	Code       string
	Name       string
	Stock      int
	PriceCents int
}

func processSale(inventory map[string]*Product, code string, quantity int) (int, error) {
	// TODO: implement the exercise.
	panic("not implemented")
}

func ex03() {
	// TODO: build an inventory, process sales, and display errors when they occur.
}
