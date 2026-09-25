package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	fmt.Println("Sending data...")
	ch <- 42 // Send data to the channel
	fmt.Println("Data sent.")
}

func receiveData(ch chan int) {
	time.Sleep(500 * time.Millisecond) // Simulate some processing time
	val := <-ch                        // Receive data from the channel
	fmt.Println("Received:", val)
}

func main() {
	// Comunication between goroutines is done using channels.
	// With buffered channels, you can send multiple values without blocking the sender until the buffer is full.
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go sendData(ch) // Start a goroutine to send data
	}

	time.Sleep(1 * time.Second) // Wait for a second to ensure all sends are done

	for i := 0; i < 5; i++ {
		go receiveData(ch) // Start a goroutine to receive data
	}

	time.Sleep(1 * time.Second) // Wait for a second to ensure all receives are done
}
