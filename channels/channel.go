package main

import (
	"fmt"
	"time"
)

func processNumber(numChan chan int) {
	// defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	for num := range numChan {
		fmt.Println("Processing number:", num) // Process the received number
		time.Sleep(time.Second) // Simulate some processing time

	}
	// fmt.Println("Processing number:", <-numChan) // Receive a number from the channel
}
func main ( ) {
	fmt.Println("Channels in Go")
	// Channels are used to communicate between goroutines.
	// They provide a way to send and receive values between goroutines.

	/* messageChan := make(chan int) // Create a channel of type int
	var waitGroup sync.WaitGroup
	waitGroup.Add(1) // Increment the WaitGroup counter
	go processNumber(messageChan, &waitGroup) // Start a goroutine that processes numbers

	messageChan <- 100 // Send a message to the channel
	waitGroup.Wait() // Wait for the goroutine to finish
	fmt.Println("Message sent to channel") */
	intChan := make(chan int) 
	go processNumber(intChan)

	var nums []int = []int { 10,20,30,40,50}
	for _,num := range nums {
		intChan <- num // Send each number to the channel
	}
	
	fmt.Println("Time Second", time.Second )
	fmt.Println("Time Second", time.Second * 2)

	/* messageChan <- "Hello, Channel!" // Send a message to the channel
	var message string = <- messageChan

	fmt.Println("Received message:", message) // Receive the message from the channel */
}	