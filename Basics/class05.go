package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo owner
}

// You can pass owner directly to the gasEngine struct as a field,
// the opwner fields becomes parts of gasEngine
//
// type gasEngine struct {
// 	mpg       uint8
// 	gallons   uint8
// 	owner
// }
// myEngine := gasEngine{mpg: 30, gallons: 10, owner{name: "John Doe"}}
// owner.name can be accessed as myEngine.name because the owner fields are promoted to gasEngine

type owner struct {
	name string
}

// This is a method of gasEngine, it can be called on any instance of gasEngine
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// Create an interface that defines a method signature for milesLeft
// Every method that has the same signature as milesLeft can be used to satisfy the engine interface
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, distance uint8) bool {
	if distance <= e.milesLeft() {
		fmt.Println("You can make it!")
		return true
	}
	fmt.Println("You cannot make it!")
	return false
}

func class05() {
	// var myEngine gasEngine the mpg and gallons fields are initialized to 0 by default (default values for uint8)
	myEngine := gasEngine{mpg: 30, gallons: 10, ownerInfo: owner{name: "John Doe"}}
	// myEngine.gallons = 11 can use this form too
	fmt.Println("My engine has", myEngine.mpg, "miles per gallon and", myEngine.gallons, "gallons of fuel.")
	fmt.Println("The owner of the engine is", myEngine.ownerInfo.name)

	// You can also create an anonymous struct without defining a new type
	// but need to initialize all fields in the struct literal
	myEngine2 := struct {
		mpg     uint8
		gallons uint8
	}{21, 14}
	fmt.Println("My engine2 has", myEngine2.mpg, "miles per gallon and", myEngine2.gallons, "gallons of fuel.")

	// Call the milesLeft method on myEngine
	fmt.Println("My engine has", myEngine.milesLeft(), "miles left.")

	// Everything that has a method called milesLeft can be used to satisfy the engine interface
	fmt.Println("Can my engine make it 300 miles?", canMakeIt(myEngine, 20))

	myElectricEngine := electricEngine{mpkwh: 3, kwh: 10, ownerInfo: owner{name: "Jane Doe"}}
	fmt.Println("Can my electric engine make it 20 miles?", canMakeIt(myElectricEngine, 20))
}
