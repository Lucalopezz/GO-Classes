package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

// O pacote constraints fornece restrições de tipo predefinidas para tipos numéricos, como constraints.Integer e constraints.Float.
// Essas restrições podem ser usadas em funções genéricas para limitar os tipos de entrada a apenas tipos numéricos.
type (
	_ constraints.Float
	_ constraints.Integer
	// tipos numéricos com sinal e sem sinal
	_ constraints.Signed
	_ constraints.Unsigned
	// num complexos
	_ constraints.Complex

	// qualquer tipo que pode ser comparado com operadores de comparação (<, <=, >, >=)
	_ constraints.Ordered
	_ cmp.Ordered // mesma coisa
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// golang.org/x/exp/constraints é um pacote que fornece restrições de tipo predefinidas para tipos numéricos,
	// como constraints.Integer e constraints.Float. Essas restrições podem ser usadas em funções genéricas para limitar os tipos de entrada a apenas tipos numéricos.

	fmt.Println("Min de inteiros:", Min(10, 20))
	fmt.Println("Min de float64:", Min(10.5, 20.5))
	fmt.Println("Min de strings:", Min("banana", "maçã"))
}
