package main

import (
	"fmt"
	"time"
)

func class10() {
	// Channels
	//
	// A channel is a typed communication mechanism used by goroutines
	// to send and receive values safely.
	//
	// Channels help coordinate goroutines and exchange data between them.
	// They are safe to use concurrently.

	// The second argument to make() is the channel buffer capacity.
	//
	// make(chan int)     -> unbuffered channel
	// make(chan int, 5)  -> buffered channel with capacity 5
	//
	// A buffered channel can hold up to 5 values before the sender blocks.
	c := make(chan int, 5)

	// Start process() in a new goroutine.
	// process() will send values into the channel.
	go process(c)

	// range receives values from the channel until the channel is closed
	// and all buffered values have been consumed.
	//
	// This loop blocks when there is no value available yet,
	// and continues when another value is sent.
	for i := range c {
		fmt.Println(i)

		// Simulate slow work by the receiver.
		time.Sleep(1 * time.Second)
	}
}

func process(c chan int) {
	// The sender owns the responsibility of closing the channel here,
	// because it knows when no more values will be sent.
	//
	// defer close(c) would also be a good option in this function,
	// because the channel should be closed when process() finishes.

	for i := 0; i < 5; i++ {
		// Send a value into the channel.
		//
		// Because this channel has capacity 5, these sends can complete
		// without waiting for the receiver, as long as the buffer is not full.
		c <- i
	}

	// Close the channel to signal:
	// "No more values will be sent."
	//
	// Closing the channel does NOT delete the buffered values.
	// The receiver can still read all values that were already sent.
	close(c)

	// This can be printed before the receiver finishes printing all values.
	//
	// Why?
	// Because process() may fill the entire buffer, close the channel,
	// and continue immediately, while class10() is consuming the buffered
	// values one by one.
	fmt.Println("Channel closed")
}
