package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var flag int64
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Escrevendo..")
		atomic.StoreInt64(&flag, 1)
	}()
}
