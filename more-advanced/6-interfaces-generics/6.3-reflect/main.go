package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = "42"

	t := reflect.TypeOf(x)

	fmt.Println("Tipo:", t)

	v := reflect.ValueOf(x)

	fmt.Println("Valor:", v)

	// Panic: reflect: chamada de reflect.Value.Int em um Value de string
	// xInt := v.Int()

	if v.CanInt() {
		xInt := v.Int()

		fmt.Println("Valor int:", xInt)
	} else {
		fmt.Println("Não é possível converter para int")
	}

	var y any = 100
	// Faz uma asserção de tipo para obter o valor subjacente
	v2 := reflect.ValueOf(&y).Elem()

	// Altera o valor de y usando reflexão
	v2.Set(reflect.ValueOf(200))

	fmt.Println("Novo valor de y:", y)
	fmt.Println("Valor de v2:", v2.Elem().Int())
}
