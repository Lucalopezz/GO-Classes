package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1               // slice2 é uma referência ao mesmo array subjacente de slice1
	slice2[0] = 10                 // modificar slice2 também modifica slice1
	fmt.Println("slice1:", slice1) // Saída: slice1: [10 2 3]

	// mas

	slice2 = append(slice2, 4)     // adicionar a slice2 pode criar um novo array subjacente se a capacidade for excedida
	slice2[1] = 20                 // modificar slice2 agora não afeta slice1
	fmt.Println("slice1:", slice1) // Saída: slice1: [10 2 3]
	fmt.Println("slice2:", slice2) // Saída: slice2: [10 20 3 4]

	// o jeito certo

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, len(slice3)) // cria um novo slice com o mesmo comprimento de slice3
	copy(slice4, slice3)               // copia o conteúdo de slice3 para slice4
	fmt.Println("slice3:", slice3)     // Saída: slice3: [1 2 3]
	slice4[0] = 10                     // modificar slice4 não afeta slice3
	fmt.Println("slice4:", slice4)     // Saída: slice4: [1 2 3]
}
