package main

import "fmt"

func class06() {
	// Pointers (initialize as nil)
	// var p *int32
	// Initialize pointer with address of an integer variable, 0 because it is the default value of int32
	var p *int32 = new(int32)
	var i int32

	// If the pointer is nil, it means it doesn't point to any valid memory address -> null pointer error
	fmt.Println("Pointer p:", p)   // prints the address of the integer variable
	fmt.Println("Value of p:", *p) // prints the value of the integer variable (0)
	*p = 42
	fmt.Println("Value of p after assignment:", *p) // prints the value of the integer variable (42)

	p = &i
	fmt.Println("Pointer p after assignment:", *p)

	*p = 100
	fmt.Println("Value of i after assignment through pointer p:", i) // prints the value of i (100)

	// Warning: when using slices, maps, and channels, they are reference types, so you don't need to use pointers
	// to modify their values. However, if you want to pass a pointer to a slice, map, or channel,
	// you can do so by using the & operator.
	slice := []int32{1, 2, 3}
	sliceCopy := slice // This creates a copy of the slice header, but both slices point to the same underlying array
	sliceCopy[0] = 100
	fmt.Println("Original slice after modifying sliceCopy:", slice) // prints [100 2 3]

	/*
		Arrays in Go are value types.

		When an array is passed to a function, the function receives a copy
		of the entire array. This means `thing` and `thing2` are two independent
		arrays stored at different memory addresses.

		Because `thing2` is a copy, modifying it inside square() does not modify
		the original `thing` array.

		For large arrays, copying the whole array may be inefficient because
		the function has to work with another copy of all its elements.

		One way to avoid copying the array is to pass a pointer:

		    func square(thing *[5]float64)

		and call it with:

		    square(&thing)

		In that case, the function receives the address of the original array
		instead of receiving a copy. Modifications made through the pointer
		will affect the original array.

		In idiomatic Go, however, slices are usually preferred over pointers
		to arrays when working with collections.
	*/

	thing := [5]float64{1, 2, 3, 4, 5}

	fmt.Println("Address of original array:", &thing)

	result := square(thing)

	fmt.Println("Address of returned array:", &result)
	fmt.Println("Original:", thing)
	fmt.Println("Squared:", result)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Println("Address of function parameter:", &thing2)

	for i := range thing2 {
		thing2[i] *= thing2[i]
	}

	return thing2
}
