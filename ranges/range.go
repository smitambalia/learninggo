package main 

import (
	"fmt"
)
func main( ) {
	fmt.Println("Ranges in Go")
	var array []int = []int {1,2,3,4,5}
	for _, item := range array {
		fmt.Println("Value: ",item)
	}
	var map1 = map[string]string {"fname" : "Smit", "lname" : "Patel" } 

	for key := range map1 {
		fmt.Println("Key: " , map1[key] )
	}
	fmt.Println(string(98))
}