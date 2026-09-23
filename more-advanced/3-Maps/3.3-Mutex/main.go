package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	// Wait for 100 goroutines to finish before exiting the program
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			mu.Lock()       // Lock the mutex to ensure exclusive access to the map
			m[i] = i
			mu.Unlock() // Unlock the mutex to allow other goroutines to access the map
			// This done make an decrement of the WaitGroup counter,
			// which is used to signal that a goroutine has finished its work.
		}()
	}

	wg.Wait() // Wait for all goroutines to finish before exiting the program

	for k, v := range m {
		fmt.Printf("%d: %d\n", k, v)
	}
}
