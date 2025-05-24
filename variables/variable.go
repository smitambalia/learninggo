package main

import "fmt"

func main() {

	var name string = "Smit"


	const (
		country string = "India"
		city    string = "Ahmedabad"
		age     int    = 25
	)

	fmt.Println("Hello, " + name + "!") // Concatenation
	fmt.Println("I am from " + city + ", " + country + ".")
	fmt.Println("I am " + fmt.Sprint(age) + " years old.") // Type conversion
}