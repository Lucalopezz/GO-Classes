package main

import (
	"fmt"
	"sync"
	"time"
)

//
// var dbData = []string{
// 	"data1",
// 	"data2",
// 	"data3",
// 	"data4",
// 	"data5",
// }

func class08() {
	sequential()
	concurrent()
}

func fakeDatabaseCall(i int) string {
	// Imagine this is:
	//
	// SELECT ...
	// HTTP request
	// reading a file
	// calling another service
	//
	// We simulate the waiting time with Sleep.
	time.Sleep(2 * time.Second)

	return dbData[i]
}

func sequential() {
	fmt.Println("=== SEQUENTIAL ===")

	start := time.Now()

	for i := 0; i < len(dbData); i++ {

		// There is NO goroutine here.
		//
		// The loop cannot continue until
		// fakeDatabaseCall finishes.
		result := fakeDatabaseCall(i)

		fmt.Println(result)
	}

	fmt.Println("Sequential time:", time.Since(start))
}

func concurrent() {
	fmt.Println("\n=== CONCURRENT ===")

	start := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < len(dbData); i++ {

		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			result := fakeDatabaseCall(index)

			fmt.Println(result)
		}(i)
	}

	// Wait for ALL database calls.
	wg.Wait()

	fmt.Println("Concurrent time:", time.Since(start))
}
