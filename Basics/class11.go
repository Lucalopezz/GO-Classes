package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	MAX_CHICKEN_PRICE float32 = 5
	MAX_TOFU_PRICE    float32 = 3
)

func class11() {
	// Unbuffered channels.
	//
	// These channels will be used by goroutines to notify us
	// when they find an acceptable chicken or tofu price.
	chickenChannel := make(chan string)
	tofuChannel := make(chan string)

	websites := []string{
		"chicken1.com",
		"chicken2.com",
		"chicken3.com",
		"chicken4.com",
	}

	for i := range websites {
		// Start one goroutine to search for a cheap chicken price
		// on this website.
		go checkChickenPrice(websites[i], chickenChannel)

		// Start another goroutine to search for a cheap tofu price
		// on this website.
		go checkTofuPrice(websites[i], tofuChannel)
	}

	// Wait until either:
	// - one chicken goroutine finds a low enough price
	// - or one tofu goroutine finds a low enough price
	//
	// The first channel that becomes ready will be selected.
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrice(web string, c chan string) {
	for {
		// Simulate checking the website once per second.
		time.Sleep(1 * time.Second)

		// Generate a random chicken price between 0 and 20.
		chickenPrice := rand.Float32() * 20

		// If the price is below our maximum acceptable price,
		// notify the receiver by sending the website through the channel.
		if chickenPrice < MAX_CHICKEN_PRICE {
			c <- web

			// Stop this goroutine after finding a valid price.
			break
		}
	}
}

func checkTofuPrice(web string, c chan string) {
	for {
		time.Sleep(1 * time.Second)

		tofuPrice := rand.Float32() * 20

		if tofuPrice < MAX_TOFU_PRICE {
			c <- web
			break
		}
	}
}

func sendMessage(chicken chan string, tofu chan string) {
	// select waits for multiple channel operations.
	//
	// It blocks until one of the cases is ready.
	// Then it executes that case.
	//
	// In this example, we react to whichever valid price
	// is found first: chicken or tofu.
	select {
	case web := <-chicken:
		fmt.Println("Chicken price is low at", web)

	case web := <-tofu:
		fmt.Println("Tofu price is low at", web)
	}
}
