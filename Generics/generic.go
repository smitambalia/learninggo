package main

import "fmt"

func printSlice[T any] (items []T) {
	for _, item := range items {
		fmt.Println(item)
	}

}


func main () {
	fmt.Println("Generics in Go")
	// Generics allow you to write functions and data structures that can work with any data type.
	// This is useful for creating reusable code that can handle different types without sacrificing type safety.

	intSlice := []int { 1,2,3,4,5 }
	printSlice(intSlice) // Calling the generic function with a slice of integers
}