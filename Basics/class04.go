package main

import (
	"fmt"
	"strings"
)

func class04() {
	// Strings, runes and bytes

	myString := "Résumé"
	fmt.Println(myString) // Résumé
	// %v value in a default format
	// %T is a Go-syntax representation of the type of the value
	fmt.Printf("%v, %T", myString[0], myString[0]) // 82, uint8
	for i, v := range myString {
		fmt.Println(i, v) // 1 233 - 3 115 - 4 115 - 5 109 - 6 233
		// Skipping the first byte of the multi-byte character
	}
	fmt.Println("The length of the string is:", len(myString)) // The length of the string is: 8

	// Convert the string to []rune when you want to work with Unicode characters.
	// can use myRunes := []rune("String") or myRunes := []rune(myString)
	myRunes := []rune(myString)

	fmt.Println("Rune count:", len(myRunes))        // 6
	fmt.Println("First rune:", string(myRunes[0]))  // R
	fmt.Println("Second rune:", string(myRunes[1])) // é

	for i, r := range myRunes {
		fmt.Printf("rune index: %d, character: %c\n", i, r)
	}

	// String concatenation
	strSlice := []string{"Hello", "World", "!"}
	catStr := ""
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// As strings are immutable, the above method is not efficient (we are creating new strings every time).
	// Instead, we can use strings.Builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr = strBuilder.String()
	fmt.Println("Concatenated string:", catStr) // Concatenated string: HelloWorld!
}
