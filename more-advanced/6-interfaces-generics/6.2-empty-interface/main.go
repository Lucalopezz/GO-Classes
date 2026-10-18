package main

import "fmt"

func main() {
	fmt.Println("String", 1, true)
	fmt.Println(1)
	fmt.Println(true)
	fmt.Println([]int{1, 2, 3, 4})

	// --------

	var (
		numero  int = 5
		numero2 any = 5
	)

	numero = numero + 1
	numero++

	// numero2 = numero2 + 1
	// numero2++

	numero2Int, ok := numero2.(int)
	if ok {
		numero2Int++
	}

	// --------

	var (
		// coloca os slices dentro de uma interface vazia, mas não é possível comparar interfaces vazias com slices
		lista  any = []int{1, 2, 3, 4, 5, 6}
		lista2 any = []int{1, 2, 3, 4, 5, 6}
	)

	// go nao consegue comparar slices, mesmo que sejam iguais, por isso o if abaixo não funciona
	if lista == lista2 {
		fmt.Println("Impossivel!")
	}

	// --------

	var x any
	x = "string"
	x = 123
	x = true

	switch v := x.(type) {
	case int:
		fmt.Println("Inteiro:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float64:", v)
	case bool:
		fmt.Println("Booleano:", v)
	}
}
