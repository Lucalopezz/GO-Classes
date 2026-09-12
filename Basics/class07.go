package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency means dealing with multiple tasks at the same time.
// The tasks can take turns running and do not need to execute simultaneously.
//
// Parallelism means executing multiple tasks at the exact same time,
// usually using multiple CPU cores.
//
// Concurrency = multiple tasks in progress.
// Parallelism = multiple tasks executing simultaneously.

var (
	// dbData represents some data that we want to fetch.
	//
	// Imagine each item requires a request to a database,
	// API, remote server, etc.
	dbData = []string{
		"data1",
		"data2",
		"data3",
		"data4",
		"data5",
	}

	// results is shared between all goroutines.
	//
	// Because multiple goroutines can access this slice,
	// we need synchronization when reading/writing it.
	results = []string{}

	// WaitGroup allows the main goroutine to wait
	// until all database calls have finished.
	wg sync.WaitGroup

	// RWMutex protects the shared "results" slice.
	//
	// Lock()  -> exclusive write access
	// RLock() -> shared read access
	m sync.RWMutex
)

func class07() {
	start := time.Now()

	for i := 0; i < len(dbData); i++ {

		// Tell the WaitGroup:
		//
		// "There is one more goroutine that must finish
		// before the program can continue."
		wg.Add(1)

		// Start dbCall concurrently.
		//
		// Without "go":
		//
		//     dbCall(i)
		//
		// every call would wait for the previous one.
		//
		// With "go":
		//
		//     go dbCall(i)
		//
		// Go schedules this function as a goroutine,
		// allowing the loop to immediately start another one.
		go dbCall(i)
	}

	// At this point all goroutines were started,
	// but they may still be running.
	//
	// Wait blocks this goroutine (main) until every
	// goroutine calls wg.Done().
	wg.Wait()

	fmt.Println()
	fmt.Println("All database calls finished.")
	fmt.Println("Total time:", time.Since(start))
	fmt.Println("Final results:", results)
}

func dbCall(i int) {
	// defer means:
	//
	// "Execute wg.Done() when this function returns."
	//
	// This is safer than putting wg.Done() manually
	// at the bottom because even if we add another
	// return later, Done() will still execute.
	defer wg.Done()

	fmt.Printf("Starting database call %d\n", i)

	// Simulate a slow database/API/network request.
	//
	// The goroutine is mostly WAITING here.
	//
	// While this goroutine waits, Go can execute
	// other goroutines.
	time.Sleep(2 * time.Second)

	// Save the result into shared memory.
	save(dbData[i])

	// Print the current state of results.
	log()

	fmt.Printf("Finished database call %d\n", i)
}

func save(result string) {
	// Lock gives this goroutine exclusive access.
	//
	// If another goroutine is currently writing,
	// this goroutine waits here.
	m.Lock()

	// defer guarantees that the mutex will be unlocked
	// when this function finishes.
	defer m.Unlock()

	results = append(results, result)
}

func log() {
	// We only need to READ results.
	//
	// RLock allows multiple readers at the same time,
	// but prevents reading while a writer has Lock().
	m.RLock()

	defer m.RUnlock()

	fmt.Println("Current results:", results)
}
