package main

import "fmt"

type mySlice []int

func main() {
	// map -> [1,2,3] -> [2,4,6]
	// filter -> [1,2,3] -> [2]
	// reduce -> [1,2,3] -> 6

	list := mySlice{1, 2, 3}

	fmt.Println("New list: ", list.Map(func(i int) int { return i * 2 }))
	fmt.Println("New list: ", list.Filter(func(i int) bool { return i%2 == 0 }))
	fmt.Println("New list: ", list.Reduce(1, func(acc int, i int) int { return acc * i }))
}

// When we see generics, we can use them to create a generic filter function that works with any type.
// However, in this case, we are using a specific type (mySlice) and a specific condition (filtering even numbers).

func (m mySlice) Filter(cond func(int) bool) mySlice {
	var result mySlice

	for _, v := range m {
		if cond(v) {
			result = append(result, v)
		}
	}

	return result
}

func (m mySlice) Map(transform func(int) int) mySlice {
	var result mySlice

	for _, v := range m {
		result = append(result, transform(v))
	}
	return result
}

func (m mySlice) Reduce(initial int, acc func(int, int) int) int {
	result := initial

	for _, v := range m {
		result = acc(result, v)
	}

	return result
}
