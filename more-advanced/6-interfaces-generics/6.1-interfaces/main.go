package main

import (
	"encoding/json"
	"fmt"
)

type Conversor interface {
	Conversor(string) (map[string]any, error)
	Formato() string
}

type ConversorJSON struct{}

type ConversorXML struct{}

func (c ConversorJSON) Conversor(data string) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal([]byte(data), &result)

	return result, err
}

func (c ConversorJSON) Formato() string {
	return "JSON"
}

type Gerencidor struct {
	conversores map[string]Conversor
}

func NewGerencidor() *Gerencidor {
	return &Gerencidor{
		conversores: make(map[string]Conversor),
	}
}

func (g *Gerencidor) RegistrarConversor(c Conversor) {
	// {
	// "json": ConversorJSON{},
	// }
	g.conversores[c.Formato()] = c
}

func (g *Gerencidor) Processar(data string, formato string) {
	conversor, ok := g.conversores[formato]
	if !ok {
		fmt.Println("Conversor não encontrado para o formato:", formato)
		return
	}

	result, err := conversor.Conversor(data)
	if err != nil {
		fmt.Println("Erro ao processar os dados:", err)
		return
	}
	fmt.Println("Resultado do processamento:", result)
}

func main() {
	gerencidor := NewGerencidor()
	gerencidor.RegistrarConversor(ConversorJSON{})
	// gerencidor.RegistrarConversor(ConversorXML{})
	dadosJSON := `{"nome": "João", "idade": 30}`
	gerencidor.Processar(dadosJSON, "JSON")
}
