package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1               // slice2 is a reference to the same underlying array as slice1
	slice2[0] = 10                 // modifying slice2 also modifies slice1
	fmt.Println("slice1:", slice1) // Output: slice1: [10 2 3]

	// but

	slice2 = append(slice2, 4)     // appending to slice2 may create a new underlying array if capacity is exceeded
	slice2[1] = 20                 // modifying slice2 now does not affect slice1
	fmt.Println("slice1:", slice1) // Output: slice1: [10 2 3]
	fmt.Println("slice2:", slice2) // Output: slice2: [10 20 3 4]

	// the right way

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, len(slice3)) // create a new slice with the same length as slice3
	copy(slice4, slice3)               // copy the contents of slice3 into slice4
	fmt.Println("slice3:", slice3)     // Output: slice3: [1 2 3]
	slice4[0] = 10                     // modifying slice4 does not affect slice3
	fmt.Println("slice4:", slice4)     // Output: slice4: [1 2 3]
}
