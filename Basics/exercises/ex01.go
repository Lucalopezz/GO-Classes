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

import "fmt"

type Expense struct {
	Description string
	Category    string
	Amount      float64
}

func summarizeExpenses(expenses []Expense) (map[string]float64, Expense, error) {
	if len(expenses) == 0 {
		err := fmt.Errorf("no expenses provided")
		return make(map[string]float64), Expense{}, err
	}
	totalByCategory := make(map[string]float64)

	highestExpense := expenses[0]

	for _, expense := range expenses {
		if expense.Amount < 0 {
			err := fmt.Errorf("negative expense amount: %v", expense.Amount)
			return make(map[string]float64), Expense{}, err
		}

		totalByCategory[expense.Category] += expense.Amount

		if expense.Amount > highestExpense.Amount {
			highestExpense = expense
		}
	}

	return totalByCategory, highestExpense, nil
}

func ex01() {
	var expenses []Expense
	expenses = append(expenses, Expense{"Groceries", "Food", 150.0})
	expenses = append(expenses, Expense{"Bus Pass", "Transportation", 50.0})
	expenses = append(expenses, Expense{"Electricity Bill", "Housing", 100.0})
	expenses = append(expenses, Expense{"Rent", "Housing", 1200.0})

	summary, highestExpense, err := summarizeExpenses(expenses)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Expense Summary:\n")
	fmt.Printf("Category\tTotal\n")
	fmt.Printf("-------------------------\n")
	fmt.Printf("Food\t\t%.2f\n", summary["Food"])
	fmt.Printf("Transportation\t%.2f\n", summary["Transportation"])
	fmt.Printf("Housing\t\t%.2f\n", summary["Housing"])
	fmt.Printf("-------------------------\n")
	fmt.Printf("Highest Expense: %s - %.2f\n", highestExpense.Description, highestExpense.Amount)
}
