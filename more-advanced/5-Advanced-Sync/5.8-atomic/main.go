package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var flag int64
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Escrevendo..")

		time.Sleep(1 * time.Second)
		// armazenando o valor 1 na variável flag de forma atômica
		// sem precisar de mutexes
		atomic.StoreInt64(&flag, 1)
		fmt.Println("Escrita concluída..")
	}()

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&flag) == 0 {
			fmt.Println("Aguardando escrita..")
			time.Sleep(100 * time.Millisecond)
		}

		fmt.Println("Leitura concluída..")
	}()

	wg.Wait()
}
