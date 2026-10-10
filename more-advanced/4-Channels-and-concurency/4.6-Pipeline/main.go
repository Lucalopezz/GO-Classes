package main

func etapa1(ch1 chan<- int, numbers []int) {
	for _, n := range numbers {
		ch1 <- n // envia o número para o canal de entrada
	}
	close(ch1) // fecha o canal de entrada quando todos os números foram enviados
}

// apenas le o ch1 e o ch2 escreve
func etapa2(ch1 <-chan int, ch2 chan<- int) {
	for n := range ch1 { // lê os números do canal de entrada
		ch2 <- n * n // envia o quadrado do número para o canal de saída
	}
	close(ch2) // fecha o canal de saída quando todos os números foram processados
}

func etapa3(ch2 <-chan int) {
	for n := range ch2 { // lê os números do canal de entrada
		println(n) // imprime o número
	}
}

func main() {
	var (
		numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // números a serem processados
		ch1     = make(chan int)
		ch2     = make(chan int)
	)

	go etapa1(ch1, numbers) // inicia a goroutine da etapa 1
	go etapa2(ch1, ch2)     // inicia a goroutine da etapa 2
	// etapas 1 e 2 já vão execultando, quando um numero é enviado para o ch1, a etapa 2 já vai processando e enviando para o ch2
	// a etapa 3 vai lendo do ch2 e imprimindo, mas como ela não é uma goroutine, o main vai esperar ela terminar
	// na forma de uma fila de produção mesmo, já vai processando e imprimindo, mas o main vai esperar ela terminar
	etapa3(ch2) // inicia a etapa 3 (não é uma goroutine, pois queremos que o main espere ela terminar)
}
