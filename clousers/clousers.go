package main

import "fmt"


func main() {
	fmt.Println("Closures in Go")
	// Example of a closure

	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	countFunc := counter()
	fmt.Println("Counter:", countFunc()) // Output: Counter: 1
	fmt.Println("Counter:", countFunc()) // Output: Counter: 2
}