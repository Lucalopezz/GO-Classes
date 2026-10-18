package conversor

import "fmt"

type (
	Conversor interface {
		Conversor(string) (map[string]any, error)
		Formato() string
	}

	Gerencidor struct {
		conversores map[string]Conversor
	}
)

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
