package main

import "fmt"

type mySlice []int

func main() {
	// map -> [1,2,3] -> [2,4,6]
	// filter -> [1,2,3] -> [2]
	// reduce -> [1,2,3] -> 6

	list := mySlice{1, 2, 3}

	fmt.Println("Nova lista: ", list.Map(func(i int) int { return i * 2 }))
	fmt.Println("Nova lista: ", list.Filter(func(i int) bool { return i%2 == 0 }))
	fmt.Println("Nova lista: ", list.Reduce(1, func(acc int, i int) int { return acc * i }))
}

// Quando virmos genéricos, poderemos usá-los para criar uma função de filtro genérica que funciona com qualquer tipo.
// Porém, neste caso, usamos um tipo específico (mySlice) e uma condição específica (filtrar números pares).

func (m mySlice) Filter(cond func(int) bool) mySlice {
	var result mySlice

	for _, v := range m {
		if cond(v) {
			result = append(result, v)
		}
	}

	return result
}

func (m mySlice) Map(transform func(int) int) mySlice {
	var result mySlice

	for _, v := range m {
		result = append(result, transform(v))
	}
	return result
}

func (m mySlice) Reduce(initial int, acc func(int, int) int) int {
	result := initial

	for _, v := range m {
		result = acc(result, v)
	}

	return result
}
