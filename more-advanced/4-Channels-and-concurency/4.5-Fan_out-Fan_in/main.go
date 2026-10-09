package main

import (
	"fmt"
	"sync"
	"time"
)

// Fan-out, Fan-in
// separa o trabalho entre várias goroutines (fan-out) e depois combina os resultados de volta em uma única goroutine (fan-in).

func produtor(numbers []int, entrada chan<- int) {
	for _, n := range numbers {
		entrada <- n // envia o número para o canal de entrada
	}
	close(entrada) // fecha o canal de entrada quando todos os números foram enviados
	// como o chanal é definido como envio apenas, fecha a entrada mas não a leitura
}

// como o channel de entrada foi fechado, a função recebe ele como leitura apenas
// o de saida é definido como envio apenas, pois a função envia o resultado para o canal de saída
func fanOut(id int, entrada <-chan int, saida chan<- int) {
	for n := range entrada { // lê os números do canal de entrada
		fmt.Printf("Goroutine %d processando número: %d\n", id, n)
		saida <- n * n // envia o quadrado do número para o canal de saída
	}

	close(saida) // fecha o canal de saída quando todos os números foram processados
}

func fanIn(saidasReadyOnly []<-chan int) <-chan int {
	var (
		saidaFinal = make(chan int)
		wg         sync.WaitGroup
	)
	// Adiciona as rotinas em um WaitGroup para esperar que todas as goroutines terminem antes de fechar o canal de saída final
	wg.Add(len(saidasReadyOnly))

	for _, saida := range saidasReadyOnly {
		// Para cada canal de saída, inicia uma goroutine que lê os números e envia para o canal de saída final
		go func() {
			defer wg.Done()

			// Transforma os canais de saida em um canal só
			for numero := range saida {
				saidaFinal <- numero
			}
		}()
	}

	// Inicia uma goroutine que espera todas as goroutines terminarem e fecha o canal de saída final
	// Isso é necessário para evitar deadlocks, pois o canal de saída final precisa ser fechado para que o loop no main() termine
	go func() {
		wg.Wait()

		// Sem esse fechamento, o loop no main() ficaria preso esperando por mais valores no canal de saída final
		close(saidaFinal)
	}()

	return saidaFinal
}

func main() {
	var (
		numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                       // números a serem processados
		entrada = make(chan int)                                             // canal de entrada
		saidas  = []chan int{make(chan int), make(chan int), make(chan int)} // canais de saída
	)

	go produtor(numbers, entrada) // inicia a goroutine do produtor

	for i, saida := range saidas {
		go fanOut(i+1, entrada, saida) // inicia as goroutines de fan-out
	}

	time.Sleep(1 * time.Second) // espera um pouco para garantir que todas as goroutines tenham terminado

	saidasReadyOnly := make([]<-chan int, len(saidas)) // cria um slice de canais de leitura apenas,
	// uma lista que vai receber os canais de saída, mas como leitura apenas, para não permitir envio de dados

	for i, saida := range saidas {
		saidasReadyOnly[i] = saida
	}

	saidaFinal := fanIn(saidasReadyOnly)

	for resultado := range saidaFinal {
		fmt.Println("Resultado recebido na saida final: ", resultado)
	}
}
