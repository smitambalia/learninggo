package main 

import (
	"fmt"
)


func main() {
	
	var arr1 [5]int = [5]int { 100, 200, 300, 400, 500 }

	arr1[0] = 250

	fmt.Println(arr1)

	var arr2 [4]bool = [4]bool {true}
	arr2[2] = true
	fmt.Println(arr2)

	array3 := [3]string {"Smit"}
	fmt.Println(array3)

	arr4 := [2][2]int {{1,2}, {3,4}}
	fmt.Println(arr4)

}