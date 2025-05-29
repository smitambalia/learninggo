package main

import "fmt"


func changeVar(num *int) {
	fmt.Println("Value before change:", *num) // Dereferencing the pointer to get the value
	*num = *num * 500 
}  

func main() {
	// Pointers in Go
	// Pointers are variables that store the memory address of another variable.
	// They are used to reference and manipulate data without copying it.

	var a int = 42
	var ptr *int = &a // ptr is a pointer to the variable a

	// Dereferencing the pointer to get the value
	value := *ptr

	// Modifying the value using the pointer
	*ptr = 100
	// Calling a function that takes a pointer as an argument
	changeVar(ptr)
	fmt.Println("Value of a:", a)            // Output: Value of a: 100
	fmt.Println("Value through pointer:", *ptr) // Output: Value through pointer: 42
	fmt.Println("Value through pointer:", value) // Output: Value through pointer: 42
}