package main

import (
	"fmt"
)
func main() {

	var marks int = 95 

	if marks >= 90  && marks <= 100 {
		fmt.Println("Grade: A")
	} else if marks >= 80{
		fmt.Println("Grade: B")	
	}

	if age := 35; age > 30 {
		fmt.Println("You are an adult")
	}
}	