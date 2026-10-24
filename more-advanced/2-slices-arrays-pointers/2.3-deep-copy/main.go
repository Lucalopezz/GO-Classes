package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 2}, {3, 4}}

	matrix2 := make([][]int, len(matrix1))
	copy(matrix2, matrix1)

	matrix2[0][0] = 100
	fmt.Println("matrix1:", matrix1) // Saída: matrix1: [[100 2] [3 4]]

	// copy faz uma cópia do slice externo, mas os slices internos ainda são compartilhados entre matrix1 e matrix2.
	fmt.Println("Cópia profunda de matrix1:", deepCopyMatrix(matrix1))
}

func deepCopyMatrix(matrix [][]int) [][]int {
	dest := make([][]int, len(matrix))

	for i, slice := range matrix {
		dest[i] = make([]int, len(slice))
		copy(dest[i], slice)
	}

	return dest
}
