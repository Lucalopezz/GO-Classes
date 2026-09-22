package main

import "fmt"

func main() {
	matrix1 := [][]int{{1, 2}, {3, 4}}

	matrix2 := make([][]int, len(matrix1))
	copy(matrix2, matrix1)

	matrix2[0][0] = 100
	fmt.Println("matrix1:", matrix1) // Output: matrix1: [[100 2] [3 4]]

	// copy makes a copy of the external slice, but the internal slices are still shared between matrix1 and matrix2.
	fmt.Println("Deep copy of matrix1:", deepCopyMatrix(matrix1))
}

func deepCopyMatrix(matrix [][]int) [][]int {
	dest := make([][]int, len(matrix))

	for i, slice := range matrix {
		dest[i] = make([]int, len(slice))
		copy(dest[i], slice)
	}

	return dest
}
