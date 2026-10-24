package main

import "fmt"

func main() {
	var (
		// Array dinâmico (slice) e array estático
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("tamanho e capacidade do slice:", len(slice), cap(slice))

	fmt.Println("tamanho e capacidade do array:", len(array), cap(array))

	// Append retorna um novo slice com os novos elementos adicionados ao final do slice original.
	// Se a capacidade do slice original não for suficiente para acomodar os novos elementos,
	// um novo array subjacente é alocado e os elementos são copiados para ele.
	slice = append(slice, 6, 7, 8)

	fmt.Println("tamanho e capacidade do slice:", len(slice), cap(slice))

	printSlice(slice)
}

// Você não pode passar um array para uma função que espera um slice
// Para passar um array a uma função que espera um slice, você pode usar o operador de slice [:] para convertê-lo em slice.
// printSlice(array[:])
func printSlice(slice []int) {
	for _, item := range slice {
		fmt.Println(item)
	}
}
