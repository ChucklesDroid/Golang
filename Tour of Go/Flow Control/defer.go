package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("World\n")
	// fmt.Println("Hello")

	fmt.Println("Counting...")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
