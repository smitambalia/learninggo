package main

import (
	"fmt"
	"sync"
)


func task(id int,w *sync.WaitGroup) {
	defer w.Done() // Decrement the counter when the goroutine completes
	fmt.Println("Task", id)
	
}
func main() {
	fmt.Println("Goroutines in Go")
	// Goroutines are lightweight threads managed by the Go runtime.
	// They allow concurrent execution of functions, making it easy to write programs that can perform multiple tasks simultaneously.

	var wg sync.WaitGroup

	for i:= 0; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go task(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks finished")
	// Note: The main function may exit before the goroutines finish executing.

}