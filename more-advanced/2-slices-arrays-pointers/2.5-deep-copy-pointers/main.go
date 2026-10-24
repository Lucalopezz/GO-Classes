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

	// A mesma lógica dos slices: se você criar person2 e atribuir person1 a ela,
	// o ponteiro para o nome será copiado, não o valor em si.

	person2 := person1

	fmt.Println("Antes da alteração:")
	fmt.Println("Pessoa 1:", *person1.Nome, person1.Age)
	fmt.Println("Pessoa 2:", *person2.Nome, person2.Age)

	fmt.Println("\nApós a alteração:")

	// após o Go 1.26; antes do Go 1.26, você pode usar a função `toPointer` para criar um ponteiro para um valor string.
	person2.Nome = new("Jane")

	fmt.Println("Pessoa 1:", *person1.Nome, person1.Age)
	fmt.Println("Pessoa 2:", *person2.Nome, person2.Age)

	person3 := deepCopyPerson(person1)

	fmt.Println("\nApós a cópia profunda:")
	fmt.Println("Pessoa 1:", *person1.Nome, person1.Age)
	person3.Nome = new("Joseph")
	fmt.Println("Pessoa 3:", *person3.Nome, person3.Age)

	// Com slices
	personsList1 := []Person{person1, person2, person3}
	personsList2 := deepCopyPersonsList(personsList1)

	personsList2[0].Nome = new("Michael")
	fmt.Println("\nApós a cópia profunda da lista de pessoas:")
	fmt.Println("Lista de pessoas 1:", *personsList1[0].Nome, personsList1[0].Age)
	fmt.Println("Lista de pessoas 2:", *personsList2[0].Nome, personsList2[0].Age)
}

// Com Go 1.26, você pode usar a função `toPointer` para criar um ponteiro para um valor string.
// Isso é útil quando você quer criar um novo ponteiro sem precisar declarar antes uma variável.
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
