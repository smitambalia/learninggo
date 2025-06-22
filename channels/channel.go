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

func sumOfTwo(resultChan chan int, num1 int, num2 int) {
	sum := num1 + num2 // Calculate the sum of two numbers
	resultChan <- sum // Send the result to the channel
}
func task(done chan bool) {
	fmt.Println("Task Started")

	done <- true // Signal that the task is done
	fmt.Println("Task Completed")
}

func sendEmail(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Signal that the email sending is done
	for email := range emailChan {
		fmt.Println("Sending email to:", email) // Simulate sending an email
  		time.Sleep(time.Second) // Simulate some delay in sending the email
 	}
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
	/* intChan := make(chan int) 
	go processNumber(intChan)

	var nums []int = []int { 10,20,30,40,50}
	for _,num := range nums {
		intChan <- num // Send each number to the channel
	}
	
	fmt.Println("Time Second", time.Second )
	fmt.Println("Time Second", time.Second * 2) */

	result := make(chan int) 

	go sumOfTwo(result, 20, 30) // Start a goroutine to calculate the sum of two numbers
	sum := <-result

	fmt.Println("Sum of two numbers:", sum) // Receive the result from the channel

	chan1 := make(chan int)
	chan2 := make(chan string)
	
	go func() {
		chan1 <- 23
	} ()

	go func() {
		chan2 <- "Smit here"
	} ()

	for i := range 2 {
		fmt.Println("Waiting for messages from channels...", i)
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num) // Receive a number from chan1
		case message := <-chan2:
			   fmt.Println("Received from chan2:", message) // Receive a message from chan2
		}
	}
	done := make(chan bool)

	// go task(done)

	emailChan := make(chan string, 100)
	go sendEmail(emailChan, done) // Start a goroutine to send emails

	for i :=range 5 {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) // Send emails to the channel
	}
	fmt.Println("Done sending emails")
	close(emailChan) // Close the channel to signal that no more emails will be sent
	<-done // Wait for the task to complete

	/* messageChan <- "Hello, Channel!" // Send a message to the channel
	var message string = <- messageChan

	fmt.Println("Received message:", message) // Receive the message from the channel */
}	