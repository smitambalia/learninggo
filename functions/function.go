package main 


import (
	"fmt"
)

func add (a int , b int) int{
	return a + b
}

func multipleData() []int {
	var data []int = []int {1, 2, 3, 4, 5}
	return data
}

func returnMap() (map[string]int) {
	var data = make(map[string]int)
	data["fname"] = 1
	data["lname"] = 2

	return data
}

func processData() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}
func main() {
	fmt.Println("Functions in Go")
	sum := add(4,5)
	fmt.Println("Sum: ", sum)

	data := multipleData()
	fmt.Println("Data: ", data)

	mapData := returnMap()
	fmt.Println("Map Data: ", mapData)

	process := processData()
	fmt.Println("Process Data: ", process(5))
}