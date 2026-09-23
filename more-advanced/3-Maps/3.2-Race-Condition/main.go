package main

import (
	"fmt"
	"time"
)

func main() {
	// Race condition occurs when multiple goroutines access a shared variable concurrently,
	// Two or more go rotines can read can access the same variable at the same time, and if
	// one of them is writing to it, it can cause a race condition.

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	fmt.Println(m)

	// Anonimous function to create a goroutine
	// 2 goroutines are created to access the same map concurrently, which can cause a race condition.
	m2 := make(map[int]int)
	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i // m2[1]
		}
	}()
	go func() {
		for i := 1000; i < 2000; i++ {
			// even though the map is being written to by one goroutine, the other goroutine
			// is trying to read from it at the same time, which can cause a race condition.
			fmt.Println(m2[i]) // m2[1000]
		}
	}()

	time.Sleep(5 * time.Second)
}
