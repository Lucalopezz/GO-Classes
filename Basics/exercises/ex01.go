/*
Exercise 01 — Expense report

Create a report for a list of monthly expenses.

 1. Calculate the total spent in each category (food, transportation, housing,
    etc.) and return these totals in a map[string]float64.
 2. Find the highest-value expense.
 3. Return an error if any expense has a negative amount.

Use a struct to represent each expense, a function with multiple return
values, slices, maps, and a loop. At the end, print a summary organized by
category and display the highest-value expense.
*/
package main

type Expense struct {
	Description string
	Category    string
	Amount      float64
}

func summarizeExpenses(expenses []Expense) (map[string]float64, Expense, error) {
	// TODO: implement the exercise.
	panic("not implemented")
}

func ex01() {
	// TODO: create test data and call summarizeExpenses.
}
