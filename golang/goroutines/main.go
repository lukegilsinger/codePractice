package main

import (
	"fmt"
)

func worker(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending:", i) // Print the value before sending
		ch <- i                    // Send value to channel
	}
	close(ch) // Close the channel when done
}

func main() {
	ch := make(chan int)

	go worker(ch) // Start the worker goroutine

	for value := range ch { // Receive values from the channel
		fmt.Println("Received:", value) // Print the received value
	}
}
