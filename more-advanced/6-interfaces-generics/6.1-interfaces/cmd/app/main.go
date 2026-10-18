package main

import (
	"fmt"

	"github.com/Lucalopezz/GO-Classes/interfaces-generics/interfaces/cmd/pkg/conversor"
	"github.com/Lucalopezz/GO-Classes/interfaces-generics/interfaces/cmd/pkg/conversor/json"
	"github.com/Lucalopezz/GO-Classes/interfaces-generics/interfaces/cmd/pkg/conversor/xml"
)

func main() {
	gerencidor := conversor.NewGerencidor()
	gerencidor.RegistrarConversor(json.ConversorJSON{})
	gerencidor.RegistrarConversor(xml.ConversorXML{})

	dadosJSON := `{"nome": "João", "idade": 30}`

	dadosXML := `<pessoa><nome>João</nome><idade>30</idade></pessoa>`

	fmt.Println("Processando dados JSON:")
	gerencidor.Processar(dadosJSON, "JSON")

	fmt.Println("\nProcessando dados XML:")
	gerencidor.Processar(dadosXML, "XML")
}
