package main

import "fmt"

func SomaInt(num []int) int {
	soma := 0
	for _, n := range num {
		soma += n
	}
	return soma
}

// Soma é uma função genérica que recebe um slice de números (int ou float64) e retorna a soma desses números.
// o int | float64 é uma restrição de tipo que permite que a função aceite apenas slices de int ou float64.
func Soma[T int | float64](num []T) T {
	var soma T
	for _, n := range num {
		soma += n
	}
	return soma
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	soma := SomaInt(numeros)
	fmt.Println("Soma dos números:", soma)

	somaGenerica := Soma(numeros)
	fmt.Println("Soma genérica dos números:", somaGenerica)
}
