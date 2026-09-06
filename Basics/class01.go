package main

import (
	"fmt"
	"unicode/utf8"
)

func class01() {
	fmt.Println("Hello World")

	var intNum int

	intNum = 10

	intNum2 := 20

	fmt.Println("Enter an integer number: ", intNum)

	fmt.Println("Enter an integer number: ", intNum2)

	// Less precise floating point numbers are float32
	// and more precise floating point numbers are float64.
	var floatNum float32 = 213213.313
	fmt.Println("Enter a floating point number: ", floatNum)

	// You need to cast types for math operations
	var result float32 = floatNum + float32(intNum)
	fmt.Println("Result of adding floatNum and intNum: ", result)

	// For division, 2 integers will result in an integer
	// so if you want a float result, you need to cast one of the integers to float
	fmt.Println("Result of dividing intNum2 by intNum: ", float32(intNum2)/float32(intNum))

	fmt.Println("Result of dividing intNum2 by intNum: ", intNum2/intNum) // 3/2 -> 1
	// If you want the remainder of a division, you can use the modulus operator %

	myString := "Hello, World!"
	myString2 := `Hello,
GO!`
	myString3 := "Hello," + "World!"

	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)

	// Like C its not the len is the byte length of the string, not the number of characters
	fmt.Println("Length of myString: ", len(myString))

	// If you want to get the number of characters in a string, you can use the utf8 package
	fmt.Println("Number of characters in myString: ", utf8.RuneCountInString(myString))

	// A rune is an alias for int32 and represents a Unicode code point
	var myRune rune = 'A'
	fmt.Println("Value of myRune: ", myRune)

	var myBool bool = true
	fmt.Println("Value of myBool: ", myBool)

	// Ways you can declare a variable
	// var myVar1 int = 10
	// myVar2 := 20
	// var1, var2 := 30, 40
	// var var myVar3 int

	// Constants are declared with the const keyword and cannot be changed after they are declared
	const myConst int = 100
	fmt.Println("Value of myConst: ", myConst)
}
