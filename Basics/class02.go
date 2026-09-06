package main

import (
	"fmt"
)

func class02() {
	// Functions and code structures
	printText := "Hello, World!"
	printMe(printText)

	numerator := 10
	denominator := 30
	result, remainder, err := intDivide(numerator, denominator)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	// } else {
	// 	fmt.Printf("%d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Println("Error:", err)
	case remainder == 0:
		fmt.Printf("The result of %d divided by %d is %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact, no remainder.\n")
	case 1, 2:
		fmt.Printf("The remainder is small: %d\n", remainder)
	default:
		fmt.Printf("The remainder is larger: %d\n", remainder)
	}
}

func printMe(print string) {
	fmt.Println(print)
}

func intDivide(a int, b int) (int, int, error) {
	var err error // default value of error is nil
	if b == 0 {
		err = fmt.Errorf("denominator cannot be zero")
		return 0, 0, err
	}

	result := a / b
	remainder := a % b
	return result, remainder, err
}
