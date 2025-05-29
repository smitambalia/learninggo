package main
import (
	"fmt"
)

func main () {
	fmt.Println("Variadic Functions in Go")
	
	// Example of a variadic function
	sum := func(numbers ...int) int {
		total := 0
		for _, number := range numbers {
			total += number
		}
		return total
	}
	var nums = []int {1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println("Sum of numbers:", result)
	// Example of a variadic function with different types

	printValues := func(values ...interface{}) {
		for _, value := range values {
			fmt.Println(value)
		}
	}
	printValues("Hello", 42, 3.14, true, "Go is fun!")
	fmt.Println("Variadic functions allow you to pass a variable number of arguments to a function.")

	var diff = []interface{}{1, "Hello", 3.14, true}
	printValues(diff...) // Using the variadic function with a slice
}