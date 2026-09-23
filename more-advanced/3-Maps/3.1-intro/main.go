package main

import "fmt"

func main() {
	// key vakue pair

	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	mapAny := map[string]any{ // JSON
		"key1": "value1",
		"key2": 2,
	}
	fmt.Println(myMap)
	fmt.Println(mapAny)

	/*
		+-----------+-------------------------+
		| Bucket 0  | [("a", 1), ("e", 5)]    |
		+-----------+-------------------------+
		| Bucket 1  | [("b", 2)]              |
		+-----------+-------------------------+
		| Bucket 2  | [("c", 3), ("f", 6)]    |
		+-----------+-------------------------+
		| Bucket 3  | [("d", 4)]              |
		+-----------+-------------------------+
	*/

	// This print can be out of order because maps are unordered collections in Go.
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// var myMap2 map[string]any // nil map
	// myMap2["key1"] = "value1" // panic: assignment to entry in nil map
	myMap2 := map[string]any{} // empty map
	myMap2["key1"] = "value1"  // works fine
	// with make
	myMap3 := make(map[string]any)
	myMap3["key1"] = "value1" // works fine

	fmt.Println(myMap2["dontExist"]) // prints <nil>
	// can cause panic if you try to access a key that doesn't exist and then try to use it as a value

	myValue, ok := myMap2["dontExist"]
	fmt.Println(myValue, ok) // prints <nil> false

	myValue2, ok2 := myMap2["key99"].(int) // type assertion, will panic if the key not exist or the value is not of type int
	if !ok2 {
		fmt.Println("Key 'key99' does not exist or is not of type int")
	} else {
		fmt.Println("Value of 'key99':", myValue2)
	}
}
