package main

import (
	"fmt"
	"os"
)

func main() {
	fs,err := os.Open("temp.txt")

	if err != nil {
		panic(err)
	}

	fileInfo,err := fs.Stat()

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("MModified time:", fileInfo.ModTime())


}