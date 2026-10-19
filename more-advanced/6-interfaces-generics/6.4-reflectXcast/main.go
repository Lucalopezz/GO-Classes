package main

import (
	"fmt"
	"reflect"
	"time"
)

func castPuro(valor any, iteracoes int) time.Duration {
	start := time.Now()

	for i := 0; i < iteracoes; i++ {
		if v, ok := valor.(int); ok {
			_ = v // simulando operação com v
		}
	}

	return time.Since(start)
}

func castComReflect(valor any, iteracoes int) time.Duration {
	start := time.Now()

	for i := 0; i < iteracoes; i++ {
		v := reflect.ValueOf(valor)

		if v.Kind() == reflect.Int {
			_ = int(v.Int()) // simulando operação com v
		}
	}

	return time.Since(start)
}

func porcentagem(t1, t2 time.Duration) float64 {
	// assume t2 < t1 (porque o cast puro sempre é mais rápido)
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {
	var (
		iteracoes          = 100_000_000
		numeroValido   any = 42
		numeroInvalido any = "não é int"
	)

	fmt.Printf("Executando comparação com %d iterações...\n", iteracoes)

	t1 := castComReflect(numeroValido, iteracoes)
	fmt.Printf("[Cast com Reflect - Número válido] Tempo total: %s\n", t1)

	t2 := castPuro(numeroValido, iteracoes)
	fmt.Printf("[Cast Puro - Número válido] Tempo total: %s\n", t2)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", porcentagem(t1, t2))

	fmt.Println("---------------------------------------")

	t3 := castComReflect(numeroInvalido, iteracoes)
	fmt.Printf("[Cast com Reflect - Número inválido] Tempo total: %s\n", t3)

	t4 := castPuro(numeroInvalido, iteracoes)
	fmt.Printf("[Cast Puro - Número inválido] Tempo total: %s\n", t4)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", porcentagem(t3, t4))
}
