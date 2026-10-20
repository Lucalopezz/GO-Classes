package main

import "fmt"

type MyInt int

type Number interface {
	// Number é uma interface que define um tipo que pode ser int ou float64.
	int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}

func main() {
	a := 10
	b := 20
	c := 10.5

	fmt.Println("Soma de inteiros:", Soma(a, b))
	fmt.Println("Soma de float64:", Soma(c, c))
	// A função Soma não pode somar diretamente um int e um float64, então precisamos converter o int para float64.
	// Ou é um outro, ou seja, a função não aceita tipos diferentes ao mesmo tempo.
	fmt.Println("Soma de int e float64:", Soma(float64(a), c)) // Conversão de int para float64

	// z1 := MyInt(5)
	// z2 := MyInt(10)
	// da erro pq MyInt não é um tipo que implementa a interface Number, então não podemos usar a função Soma com MyInt.
	// fmt.Println("Soma de MyInt:", Soma(z1, z2))
}
