package main

import "fmt"

func class03() {
	// Arrays, slices, maps and loops
	intArray := [5]int{1, 2, 3, 4, 5}
	var intArray2 [3]int
	fmt.Println(intArray[0])  // 1
	fmt.Println(intArray2[0]) // 0

	// Memory address of array
	fmt.Println(&intArray[0]) // 0xc00000e030
	// Contiguous memory allocation

	// Slices
	// The slice is a reference type, it does not store any data,
	// it just describes a section of an underlying array.

	// Slices wrap arrays to give a more general, powerful, and convenient
	// interface to sequences of data.

	intSlice := []int32{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	// Before appending, the length and capacity of the slice is 5
	fmt.Printf("The length of the slice is: %v, and the capacity is: %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 6)
	fmt.Println(intSlice)
	// After appending, the length is 6 and capacity of the slice is 12
	fmt.Printf("The length of the slice is: %v, and the capacity is: %v", len(intSlice), cap(intSlice))

	intSlice2 := []int32{1, 2, 3, 4, 5}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\n The union of the two slices is: %v", intSlice)

	slice3 := make([]int32, 5, 10) // Slice with length 5 and capacity 10
	fmt.Printf("\n The length of the slice is: %v, and the capacity is: %v", len(slice3), cap(slice3))

	// Maps
	// Key-value pairs, unordered collection of data
	// {"key1": "value1", "key2": "value2"}

	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	myMap["key1"] = 1
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"key1": 1, "key2": 2}
	fmt.Println(myMap2)

	// If the key does not exist, the zero value of the value type is returned
	// because 0 is default value of uint8
	// There is this second way to check if the key exists in the map, by using the "comma ok" idiom
	value, ok := myMap2["key3"]
	if !ok {
		fmt.Println("Key does not exist, value is:", value) // Key does not exist, value is: 0
	}
	// delete(myMap2, "key1")
	// fmt.Println(myMap2) // map[key2:2]

	// For loop
	for key := range myMap2 {
		fmt.Println("Key:", key, "Value:", myMap2[key])
	}
	for key, value := range myMap2 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// While loop in go is implemented using for loop
	i := 0
	for i < 5 {
		// You can use the "break" statement to exit the loop
		fmt.Println("i is:", i)
		i++
	}

	// The normal use
	for i := 0; i < 5; i++ {
		fmt.Println("i is:", i)
	}
	// i++, i--, i += 2, i -= 2, i *= 2, i /= 2 are all valid increment/decrement operations in Go

	// For general use, always pre alocate the memory for slices and maps,
	// its more efficient than appending to them, because appending to a slice or map may cause the underlying array
	// to be reallocated and copied to a new location in memory, which can be expensive in terms of performance.
}
