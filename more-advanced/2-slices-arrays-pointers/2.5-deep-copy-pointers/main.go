package main

import "fmt"

type Person struct {
	Nome *string
	Age  int
}

func main() {
	var name string = "John"
	person1 := Person{
		Nome: &name,
		Age:  30,
	}

	// Same logic with slices, if you create a person2 and assign person1 to it,
	// it will copy the pointer to the name, not the value itself.

	person2 := person1

	fmt.Println("Before change:")
	fmt.Println("Person 1:", *person1.Nome, person1.Age)
	fmt.Println("Person 2:", *person2.Nome, person2.Age)

	fmt.Println("\nAfter change:")

	// afeter go 1.26, before go 1.26, you can use the new function `toPointer` to create a pointer to a string value.
	person2.Nome = new("Jane")

	fmt.Println("Person 1:", *person1.Nome, person1.Age)
	fmt.Println("Person 2:", *person2.Nome, person2.Age)

	person3 := deepCopyPerson(person1)

	fmt.Println("\nAfter deep copy:")
	fmt.Println("Person 1:", *person1.Nome, person1.Age)
	person3.Nome = new("Joseph")
	fmt.Println("Person 3:", *person3.Nome, person3.Age)

	// With slices
	personsList1 := []Person{person1, person2, person3}
	personsList2 := deepCopyPersonsList(personsList1)

	personsList2[0].Nome = new("Michael")
	fmt.Println("\nAfter deep copy of persons list:")
	fmt.Println("Persons List 1:", *personsList1[0].Nome, personsList1[0].Age)
	fmt.Println("Persons List 2:", *personsList2[0].Nome, personsList2[0].Age)
}

// With go 1.26, you can use the new function `toPointer` to create a pointer to a string value.
// This is useful when you want to create a new pointer without having to declare a variable first.
//
// func toPointer(s string) *string {
// 	return &s
// }

func deepCopyPerson(p Person) Person {
	var dest Person

	dest.Nome = new(*p.Nome)
	dest.Age = p.Age

	return dest
}

func deepCopyPersonsList(persons []Person) []Person {
	dest := make([]Person, len(persons))

	for i, p := range persons {
		dest[i] = deepCopyPerson(p)
	}

	return dest
}
