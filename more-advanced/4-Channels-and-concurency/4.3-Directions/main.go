package main

import (
	"fmt"
	"time"
)

// Can only send data to the channel chan<-
func sendData(ch chan<- int) {
	fmt.Println("Sending data...")
	ch <- 42
	fmt.Println("Data sent.")
}

// Can only receive data from the channel <-chan
func receiveData(ch <-chan int) {
	time.Sleep(500 * time.Millisecond)
	val := <-ch
	fmt.Println("Received:", val)
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go sendData(ch)
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 5; i++ {
		go receiveData(ch) // Start a goroutine to receive data
	}

	time.Sleep(1 * time.Second) // Wait for a second to ensure all receives are done
}
