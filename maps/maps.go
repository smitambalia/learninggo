package main

import (
	"fmt"
	"maps"
)
func main() {
	fmt.Println("Maps in Go")
	// Declaring a map with string keys and string values
	var m1 = make(map[string]string) // 
	m1["firstName"] = "Smit"
	m1["lastName"] = "Patel"

	fmt.Println(m1["firstName"]) // Output: Smit
	fmt.Println(m1["lastName"]) // Output: Patel

	var m2 = make(map[string]int) 

	m2["age"] = 34
	m2["height"] = 176

	fmt.Println(m2)

	// delete(m2,"age")
	fmt.Println(m2)
	// clear(m2)
	fmt.Println(m2)

	m3 := make(map[string]int)

	m3["age"] = 34


	_, isExist := m2["age"]
	if isExist {
		fmt.Printf("Key 'age' exists in the map\n")
	} else {
		fmt.Printf("Key 'age' does not exist in the map %v \n", isExist)
	}

	fmt.Println(maps.Equal(m3,m2))
}	