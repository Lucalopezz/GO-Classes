package main

import "fmt"

func main() {
	var (
		// Dynamic array (slice) and static array
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("size and capacity of slice:", len(slice), cap(slice))

	fmt.Println("size and capacity of array:", len(array), cap(array))

	// Append returns a new slice with the new elements added to the end of the original slice.
	// If the capacity of the original slice is not enough to accommodate the new elements,
	// a new underlying array is allocated and the elements are copied over.
	slice = append(slice, 6, 7, 8)

	fmt.Println("size and capacity of slice:", len(slice), cap(slice))

	printSlice(slice)
}

// You cant pass an array to a function that expects a slice
// For pass an array to a function that expects a slice, you can use the slice operator [:] to convert the array to a slice.
// printSlice(array[:])
func printSlice(slice []int) {
	for _, item := range slice {
		fmt.Println(item)
	}
}
