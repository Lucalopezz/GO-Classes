package main

import (
	"fmt"
	"time"
)

func main() {
	// Uma condição de corrida ocorre quando várias goroutines acessam uma variável compartilhada concorrentemente,
	// duas ou mais goroutines podem acessar a mesma variável ao mesmo tempo e, se
	// uma delas estiver escrevendo nela, isso pode causar uma condição de corrida.

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	fmt.Println(m)

	// Função anônima para criar uma goroutine
	// Duas goroutines são criadas para acessar o mesmo map concorrentemente, o que pode causar uma condição de corrida.
	m2 := make(map[int]int)
	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i // m2[1]
		}
	}()
	go func() {
		for i := 1000; i < 2000; i++ {
			// embora uma goroutine esteja escrevendo no map, a outra goroutine
			// tenta lê-lo ao mesmo tempo, o que pode causar uma condição de corrida.
			fmt.Println(m2[i]) // m2[1000]
		}
	}()

	time.Sleep(5 * time.Second)
}
