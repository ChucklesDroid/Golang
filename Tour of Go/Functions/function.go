package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Addition of 11 and 13 is", add(11, 13))
}
