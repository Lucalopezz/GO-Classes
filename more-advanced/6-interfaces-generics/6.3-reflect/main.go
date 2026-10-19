package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = "42"

	t := reflect.TypeOf(x)

	fmt.Println("Type:", t)

	v := reflect.ValueOf(x)

	fmt.Println("Value:", v)

	// Panic: reflect: call of reflect.Value.Int on string Value
	// xInt := v.Int()

	if v.CanInt() {
		xInt := v.Int()

		fmt.Println("Int value:", xInt)
	} else {
		fmt.Println("Cannot convert to int")
	}

	var y any = 100
	// Perform type assertion to get the underlying value
	v2 := reflect.ValueOf(&y).Elem()

	// Alter the value of y using reflection
	v2.Set(reflect.ValueOf(200))

	fmt.Println("New value of y:", y)
	fmt.Println("Value of v2:", v2.Elem().Int())
}
