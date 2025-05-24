package main

import (
	"fmt"
	"time"
)

func main() {

	var age int = 35

	switch age {
		case 18:
			fmt.Println("You are a teenager")
		case 25:
			fmt.Println("You are in your twenties")
		case 30,35:
			fmt.Println("You are an adult")
		default:
			fmt.Println("You are not a teenager, in your twenties, or an adult")
	}

	fmt.Printf("time : %v, %T \n" , time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("It's a weekday.")

	}
}