package main

import (
	"fmt"
	"time"
)

// var wg = sync.WaitGroup{}
//
//	var dbData = []string{
//		"id1",
//		"id2",
//		"id3",
//		"id4",
//		"id5",
//	}
func class09() {
	t0 := time.Now()

	// Create 10,000 goroutines.
	//
	// Each iteration:
	// 1. Adds one task to the WaitGroup.
	// 2. Starts dbCall as a goroutine.
	//
	// These goroutines run concurrently.
	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go dbCall2(i)
	}

	// Block the main goroutine until all 10,000
	// goroutines call wg.Done().
	wg.Wait()

	fmt.Printf(
		"\nTotal execution time: %v\n",
		time.Since(t0),
	)
}

func dbCall2(i int) {
	// Simulate a slow database/API call.
	//
	// This goroutine waits for 2 seconds.
	// While it is waiting, other goroutines can run.
	delay := 2000

	time.Sleep(
		time.Duration(delay) * time.Millisecond,
	)

	// Tell the WaitGroup that this goroutine finished.
	wg.Done()
}
