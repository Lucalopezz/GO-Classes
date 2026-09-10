/*
Exercise 05 — Cart checkout with discount policies

Implement a shopping cart checkout using integer cents to avoid precision
problems with float64.

 1. Calculate the subtotal of all items.
 2. Validate that the price and quantity are not negative and that the
    quantity is greater than zero.
 3. Apply a discount policy received through an interface.
 4. Never allow the final total to become negative.

Create at least two concrete policies, such as a percentage discount and a
fixed-value coupon, implementing the method required by the interface. Test
the checkout with each policy and also without a discount. Return errors for
an invalid cart and use methods whenever they make the code clearer.
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
	// TODO: implement the exercise.
	panic("not implemented")
}

func ex05() {
	// TODO: create items, implement discount policies, and test the checkout.
}
