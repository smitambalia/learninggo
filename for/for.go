package main

import (
	"fmt"
)
func main() {

	for i :=0; i < 10; i++ {
		fmt.Println("i : ", i)
	}

	var j int = 0

	for j < 3 {
		fmt.Println("j : ", j)
		j++
	} 
}