package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Println("Value of a:", a)             // Output: Value of a: 10
	fmt.Println("Address of a:", &a)          // Output: Address of a: 0xc0000140b0 (example address)
	fmt.Println("Value of b:", b)             // Output: Value of b: 0xc0000140b0 (example address)
	fmt.Println("Value pointed to by b:", *b) // Output: Value pointed to by b: 10

	p := NewPerson("Alice", 30)
	fmt.Printf("Person: %+v, Phone: %s\n", p, p.Phone()) // Output: Person: {Name:Alice Age:30 phone:<nil>}, Phone:
}

// Capitalize the first letter of a struct field to make it exported and accessible from other packages.
// Unexported fields (starting with a lowercase letter) are only accessible within the same package.

type Person struct {
	Name  string
	Age   int
	phone *string
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) Phone() string {
	if p.phone == nil {
		return ""
	}
	return *p.phone
}

func (p *Person) UpdateAge(newAge int) {
	// Note: We are using a pointer receiver here, so we can modify the original struct.
	// If we used a value receiver (p Person), we would be modifying a copy of the struct, and the original struct would remain unchanged.
	p.Age = newAge
}

func (p *Person) UpdatePhone(phone string) {
	p.phone = &phone
}
