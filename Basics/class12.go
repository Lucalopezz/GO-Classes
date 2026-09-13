package main

import "fmt"

func class12() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of intSlice:", sumSlice[int](intSlice))

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of floatSlice:", sumSlice[float64](floatSlice))

	fmt.Println("Is intSlice empty?", isEmpty(intSlice)) // compiler can infer the type here
	// but for example, we can explicitly specify the type
	// from a json, so we need to specify the type explicitly
	//
	// Json function
	// json.Unmarshal([]byte(`{"name": "John", "age": 30}`), &data)
	// json string to struct, we need to specify the type explicitly
}

// Generic syntax
// [T type | type2 ]
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Example of an any type generic function
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// With structs, we can use generics to create a generic struct that can hold any type of data. For example,
// we can create a generic struct called "Box" that can hold any type of data.

type gasEngine struct {
	gallons int
	mpg     float64
}
type electricEngine struct {
	kwh  int
	mpkw float64
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}
