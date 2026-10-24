package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Println("Valor de a:", a)            // Saída: Valor de a: 10
	fmt.Println("Endereço de a:", &a)        // Saída: Endereço de a: 0xc0000140b0 (endereço de exemplo)
	fmt.Println("Valor de b:", b)            // Saída: Valor de b: 0xc0000140b0 (endereço de exemplo)
	fmt.Println("Valor apontado por b:", *b) // Saída: Valor apontado por b: 10

	p := NewPerson("Alice", 30)
	fmt.Printf("Pessoa: %+v, Telefone: %s\n", p, p.Phone()) // Saída: Pessoa: {Name:Alice Age:30 phone:<nil>}, Telefone:
}

// Use a primeira letra maiúscula em um campo da struct para exportá-lo e torná-lo acessível a outros pacotes.
// Campos não exportados (que começam com letra minúscula) só podem ser acessados dentro do mesmo pacote.

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
	// Observação: usamos um receptor ponteiro aqui para poder modificar a struct original.
	// Se usássemos um receptor por valor (p Person), modificaríamos uma cópia da struct e a original permaneceria inalterada.
	p.Age = newAge
}

func (p *Person) UpdatePhone(phone string) {
	p.phone = &phone
}
