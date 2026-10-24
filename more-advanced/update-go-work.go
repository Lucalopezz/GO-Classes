// Como isto é um script, a função main pode gerar um aviso em algumas IDEs, mas é um programa Go válido.
// A função main é o ponto de entrada do programa, onde a execução começa. Neste caso,
// a função main atualiza o arquivo go.work com todos os módulos encontrados no diretório atual e em seus subdiretórios.
package main

// Para compilar um script sem módulo, use este comando
// GO111MODULE=off go build -o update-go-work update-go-work.go

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var dirs []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "go.mod" {
			dirs = append(dirs, filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		fmt.Println("Erro ao percorrer o caminho:", err)
		os.Exit(1)
	}

	if len(dirs) == 0 {
		fmt.Println("Nenhum arquivo go.mod encontrado no diretório atual ou em seus subdiretórios.")
	}

	// Decide entre init e use
	var cmd *exec.Cmd

	if _, err := os.Stat("go.work"); os.IsNotExist(err) {
		// Se não existir, inicializa o espaço de trabalho com todos os módulos
		args := append([]string{"work", "init"}, dirs...)
		cmd = exec.Command("go", args...) // go work init <dirs...>
		fmt.Println("Atualizando go.work: %v\n", args)
	} else {
		// Se já existir, usa o espaço de trabalho com todos os módulos
		args := append([]string{"work", "use"}, dirs...)
		cmd = exec.Command("go", args...) // go work use <dirs...>
		fmt.Println("Atualizando go.work: %v\n", args)
	}

	// A saída e o erro padrão do comando são direcionados para a saída e o erro padrão do programa.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Go permite declarar uma variável antes da condição de um if.
	// A variável existe apenas no escopo do if/else.
	//
	// Aqui, cmd.Run() é executado e o erro retornado é armazenado em err.
	// Se err não for nil, o comando falhou e o bloco de tratamento de erro é executado.
	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao executar o comando go:", err)
		os.Exit(1)
	}

	fmt.Println("Arquivo go.work atualizado com sucesso.")
}
