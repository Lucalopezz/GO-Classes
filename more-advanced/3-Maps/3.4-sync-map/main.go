package main

// Sync.Map is a concurrent map implementation in Go that provides safe access to a map from multiple goroutines without the need for explicit locking.
// It is part of the sync package and is designed for scenarios where multiple goroutines need to read and write to a shared map concurrently.

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map // Use sync.Map for concurrent access to the map
		wg sync.WaitGroup
	)

	// sync.Map itnot the same type as a regular map, so we cannot use the same syntax to add key-value pairs to the map.
	// Instead, we use the Store method of sync.Map to add a key-value pair to the map.

	// Wait for 100 goroutines to finish before exiting the program
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			// Make the same thing as the the mutex
			m.Store(i, i) // Use the Store method of sync.Map to add a key-value pair to the map
		}()
	}

	wg.Wait() // Wait for all goroutines to finish before exiting the program

	// For does not work with sync.Map, so we use the Range method to iterate over the map
	m.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true // Continue iterating over the map
	})

	value, ok := m.Load(50)
	if ok {
		fmt.Printf("Value for key 50: %d\n", value)
	} else {
		fmt.Println("Key 50 not found")
	}
}
